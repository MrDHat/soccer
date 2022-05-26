package repository

import (
	"context"
	"database/sql"
	"time"

	"soccer-manager/db"
	"soccer-manager/db/models"
	"soccer-manager/logger"
)

type UserRepo interface {
	SaveWithTeamAndPlayers(ctx context.Context, user *models.User, team *models.Team) error
}

type userRepo struct {
	dbInstance db.DBInstance
}

func (repo *userRepo) SaveWithTeamAndPlayers(ctx context.Context, user *models.User, team *models.Team) error {
	var (
		groupError = "SAVE_USER_WITH_TEAM_AND_PLAYERS"
		db         = repo.dbInstance.GetWritableDB()
		nowTime    = time.Now().Unix()
	)

	logger.Log.Info("begin transaction for saving user with team & players")
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

	logger.Log.Info("saving team")
	team.CreatedAt = &nowTime
	team.UpdatedAt = &nowTime
	_, err = db.Insert(team)
	if err != nil {
		rErr := db.Rollback()
		if rErr != nil {
			logger.Log.WithError(rErr).Error(groupError)
			return rErr
		}
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	logger.Log.Info("done saving team")

	logger.Log.Info("saving user")
	user.Team = team
	user.CreatedAt = &nowTime
	user.UpdatedAt = &nowTime
	_, err = db.Insert(user)
	if err != nil {
		rErr := db.Rollback()
		if rErr != nil {
			logger.Log.WithError(rErr).Error(groupError)
			return rErr
		}
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	logger.Log.Info("done saving user")

	logger.Log.Info("saving players")
	for _, player := range team.Players {
		player.CreatedAt = &nowTime
		player.UpdatedAt = &nowTime
		player.Team = team
	}
	_, err = db.InsertMulti(len(team.Players), team.Players)
	if err != nil {
		rErr := db.Rollback()
		if rErr != nil {
			logger.Log.WithError(rErr).Error(groupError)
			return rErr
		}
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	logger.Log.Info("done saving players")

	logger.Log.Info("committing transaction")
	err = db.Commit()
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	logger.Log.Info("transaction done")

	return nil
}

func NewUserRepo(
	dbInstance db.DBInstance,
) UserRepo {
	return &userRepo{
		dbInstance: dbInstance,
	}
}
