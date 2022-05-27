package repository

import (
	"context"

	"soccer-manager/constants"
	"soccer-manager/db"
	"soccer-manager/db/models"
	"soccer-manager/logger"

	"github.com/thoas/go-funk"
)

type PlayerRepo interface {
	FindOne(ctx context.Context, query models.PlayerQuery) (*models.Player, error)
	FindAll(ctx context.Context, query models.PlayerQuery, fetchRelated bool, returnCount bool) ([]*models.Player, int64, error)
}

type playerRepo struct {
	dbInstance db.DBInstance
}

func (repo *playerRepo) FindOne(ctx context.Context, query models.PlayerQuery) (*models.Player, error) {
	var (
		groupError = "FIND_ONE_PLAYER"
		db         = repo.dbInstance.GetReadableDB()
		player     = &models.Player{}
	)

	// failsafe for empty query
	if funk.IsEmpty(query) {
		return nil, nil
	}

	qs := db.QueryTable(player)

	if query.ID != 0 {
		qs = qs.Filter("id", query.ID)
	}
	if query.Team != nil && query.Team.ID != 0 {
		qs = qs.Filter("team_id", query.Team.ID)
	}

	err := qs.One(player)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return nil, err
	}

	return player, nil
}

func (repo *playerRepo) FindAll(ctx context.Context, query models.PlayerQuery, fetchRelated bool, returnCount bool) ([]*models.Player, int64, error) {
	groupError := "FIND_ALL_PLAYERS"
	var res []*models.Player

	sortingOrder := "-"
	if query.SortOrder != nil && *query.SortOrder == constants.SortOrderAscending {
		sortingOrder = ""
	}
	orderBy := "id"
	if query.SortBy != nil {
		orderBy = *query.SortBy
	}

	qs := repo.dbInstance.GetReadableDB().QueryTable(new(models.Player))
	qs = qs.OrderBy(sortingOrder + orderBy)

	if query.Team != nil && query.Team.ID != 0 {
		qs = qs.Filter("team_id", query.Team.ID)
	}
	if query.PlayerType != "" {
		qs = qs.Filter("player_type", query.PlayerType)
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

func NewPlayerRepo(
	dbInstance db.DBInstance,
) PlayerRepo {
	return &playerRepo{
		dbInstance: dbInstance,
	}
}
