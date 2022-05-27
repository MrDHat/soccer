package repository

import (
	"context"

	"soccer-manager/db"
	"soccer-manager/db/models"
	"soccer-manager/logger"

	"github.com/thoas/go-funk"
)

type TeamRepo interface {
	FindOne(ctx context.Context, query models.TeamQuery) (*models.Team, error)
}

type teamRepo struct {
	dbInstance db.DBInstance
}

func (repo *teamRepo) FindOne(ctx context.Context, query models.TeamQuery) (*models.Team, error) {
	var (
		groupError = "FIND_ONE_TEAM"
		db         = repo.dbInstance.GetReadableDB()
		team       = &models.Team{}
	)

	// failsafe for empty query
	if funk.IsEmpty(query) {
		return nil, nil
	}

	qs := db.QueryTable(team)

	if query.ID != 0 {
		qs = qs.Filter("id", query.ID)
	}

	err := qs.One(team)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return nil, err
	}

	return team, nil
}

func NewTeamRepo(
	dbInstance db.DBInstance,
) TeamRepo {
	return &teamRepo{
		dbInstance: dbInstance,
	}
}
