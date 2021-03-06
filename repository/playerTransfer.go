package repository

import (
	"context"
	"database/sql"
	"time"

	"soccer-manager/constants"
	"soccer-manager/db"
	"soccer-manager/db/models"
	"soccer-manager/logger"

	"github.com/astaxie/beego/orm"
	"github.com/thoas/go-funk"
)

type PlayerTransferRepo interface {
	Create(ctx context.Context, playerTransfer *models.PlayerTransfer) error
	FindOne(ctx context.Context, query models.PlayerTransferQuery, returnRelated bool) (*models.PlayerTransfer, error)
	FindAll(ctx context.Context, query models.PlayerTransferQuery, fetchRelated bool, returnCount bool) ([]*models.PlayerTransfer, int64, error)
	CompleteTransfer(ctx context.Context, doc *models.PlayerTransfer, newTeamID int64, newPlayerValue int64) error
}

type playerTransferRepo struct {
	db db.DBInstance
}

func (repo *playerTransferRepo) Create(ctx context.Context, playerTransfer *models.PlayerTransfer) error {
	var (
		groupError = "CREATE_PLAYER_TRANSFER"
		db         = repo.db.GetWritableDB()
		nowTime    = time.Now().Unix()
	)

	logger.Log.Info("begin transaction for saving player transfer with player")
	err := db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: false,
	})
	if err != nil {
		rErr := db.Rollback()
		if rErr != nil {
			logger.Log.WithError(err).Error(groupError)
			return err
		}
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	logger.Log.Info("saving player transfer")
	playerTransfer.CreatedAt = &nowTime
	playerTransfer.UpdatedAt = &nowTime
	_, err = db.Insert(playerTransfer)
	if err != nil {
		rErr := db.Rollback()
		if rErr != nil {
			logger.Log.WithError(rErr).Error(groupError)
			return rErr
		}
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	logger.Log.Info("updating player")
	p := &models.Player{
		Base: models.Base{
			ID:        playerTransfer.Player.ID,
			UpdatedAt: &nowTime,
		},
		TransferStatus: string(constants.PlayerTransferStatusOnSale),
	}
	_, err = db.Update(p, []string{"updated_at", "transfer_status"}...)
	if err != nil {
		rErr := db.Rollback()
		if rErr != nil {
			logger.Log.WithError(rErr).Error(groupError)
			return rErr
		}
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	logger.Log.Info("committing transaction")
	err = db.Commit()
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	logger.Log.Info("transaction done")

	return err
}

func (repo *playerTransferRepo) FindOne(ctx context.Context, query models.PlayerTransferQuery, returnRelated bool) (*models.PlayerTransfer, error) {
	var (
		groupError = "FIND_ONE_PLAYER_TRANSFER"
		db         = repo.db.GetReadableDB()
		transfer   = &models.PlayerTransfer{}
	)

	// failsafe for empty query
	if funk.IsEmpty(query) {
		return nil, nil
	}

	qs := db.QueryTable(transfer)

	if query.ID != 0 {
		qs = qs.Filter("id", query.ID)
	}
	if query.OwnerTeam != nil && query.OwnerTeam.ID != 0 {
		qs = qs.Filter("owner_team_id", query.OwnerTeam.ID)
	}
	if query.Player != nil && query.Player.ID != 0 {
		qs = qs.Filter("player_id", query.Player.ID)
	}

	if returnRelated {
		qs = qs.RelatedSel()
	}

	err := qs.One(transfer)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return nil, err
	}

	return transfer, nil
}
func (repo *playerTransferRepo) FindAll(ctx context.Context, query models.PlayerTransferQuery, fetchRelated bool, returnCount bool) ([]*models.PlayerTransfer, int64, error) {
	groupError := "FIND_ALL_PLAYER_TRANSFERS"
	var res []*models.PlayerTransfer

	sortingOrder := "-"
	if query.SortOrder != nil && *query.SortOrder == constants.SortOrderAscending {
		sortingOrder = ""
	}
	orderBy := "id"
	if query.SortBy != nil {
		orderBy = *query.SortBy
	}

	qs := repo.db.GetReadableDB().QueryTable(new(models.PlayerTransfer))
	qs = qs.OrderBy(sortingOrder + orderBy)

	if query.OwnerTeam != nil && query.OwnerTeam.ID != 0 {
		qs = qs.Filter("owner_team_id", query.OwnerTeam.ID)
	}
	if query.NotInTeamID != nil && *query.NotInTeamID != 0 {
		qs = qs.Exclude("owner_team_id", query.OwnerTeam.ID)
	}
	if query.Status != "" {
		qs = qs.Filter("status", query.Status)
	}

	countQs := qs

	if query.Page != nil && query.Limit != nil {
		qs = qs.Offset((*query.Page - 1) * *query.Limit).Limit(*query.Limit)
	}

	if fetchRelated {
		qs = qs.RelatedSel()
	}
	_, err := qs.All(&res)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return res, 0, err
	}

	count := int64(0)
	if returnCount {
		count, err = countQs.Count()
		if err != nil {
			logger.Log.WithError(err).Error(groupError)
			return res, 0, err
		}
	}

	return res, count, nil
}

func (repo *playerTransferRepo) CompleteTransfer(ctx context.Context, doc *models.PlayerTransfer, newTeamID int64, newPlayerValue int64) error {
	var (
		groupError = "COMPLETE_PLAYER_TRANSFER"
		db         = repo.db.GetWritableDB()
		nowTime    = time.Now().Unix()
		amount     = doc.AmountInDollars
		buyerTeam  = models.Team{
			Base: models.Base{
				ID: newTeamID,
			},
		}
		sellerTeam = doc.OwnerTeam
		player     = doc.Player
	)

	logger.Log.Info("begin transaction for completing player transfer")
	err := db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: false,
	})
	if err != nil {
		rErr := db.Rollback()
		if rErr != nil {
			logger.Log.WithError(err).Error(groupError)
			return err
		}
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	logger.Log.Info("deduct amount from the buyer team")
	_, err = db.QueryTable(buyerTeam).Filter("id", buyerTeam.ID).Update(orm.Params{
		"updated_at":                  nowTime,
		"remaining_budget_in_dollars": orm.ColValue(orm.ColMinus, amount),
	})
	if err != nil {
		rErr := db.Rollback()
		if rErr != nil {
			logger.Log.WithError(err).Error(groupError)
			return err
		}
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	logger.Log.Info("transfer amount to the seller team")
	_, err = db.QueryTable(sellerTeam).Filter("id", sellerTeam.ID).Update(orm.Params{
		"updated_at":                  nowTime,
		"remaining_budget_in_dollars": orm.ColValue(orm.ColAdd, amount),
	})
	if err != nil {
		rErr := db.Rollback()
		if rErr != nil {
			logger.Log.WithError(err).Error(groupError)
			return err
		}
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	logger.Log.Info("update player transfer status")
	_, err = db.QueryTable(doc).Filter("id", doc.ID).Update(orm.Params{
		"updated_at":   nowTime,
		"status":       string(constants.TransferStatusComplete),
		"completed_at": nowTime,
	})
	if err != nil {
		rErr := db.Rollback()
		if rErr != nil {
			logger.Log.WithError(err).Error(groupError)
			return err
		}
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	logger.Log.Info("update player data")
	_, err = db.QueryTable(player).Filter("id", player.ID).Update(orm.Params{
		"updated_at":               nowTime,
		"team_id":                  newTeamID,
		"transfer_status":          string(constants.PlayerTransferStatusOwned),
		"current_value_in_dollars": newPlayerValue,
	})
	if err != nil {
		rErr := db.Rollback()
		if rErr != nil {
			logger.Log.WithError(err).Error(groupError)
			return err
		}
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	logger.Log.Info("committing transaction")
	err = db.Commit()
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	logger.Log.Info("transaction done")

	return nil
}

func NewPlayerTransferRepo(db db.DBInstance) PlayerTransferRepo {
	return &playerTransferRepo{
		db: db,
	}
}
