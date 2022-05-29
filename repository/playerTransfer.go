package repository

import (
	"context"
	"database/sql"
	"time"

	"soccer-manager/constants"
	"soccer-manager/db"
	"soccer-manager/db/models"
	"soccer-manager/logger"
)

type PlayerTransferRepo interface {
	Create(ctx context.Context, playerTransfer *models.PlayerTransfer) error
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

func NewPlayerTransferRepo(db db.DBInstance) PlayerTransferRepo {
	return &playerTransferRepo{
		db: db,
	}
}
