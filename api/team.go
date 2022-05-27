package api

import (
	"context"
	"errors"

	"soccer-manager/api/helpers"
	apiutils "soccer-manager/api/utils"
	"soccer-manager/constants"
	"soccer-manager/db/models"
	graphmodel "soccer-manager/graph/model"
	"soccer-manager/logger"
	"soccer-manager/repository"

	"github.com/astaxie/beego/orm"
)

type Team interface {
	My(ctx context.Context) (*graphmodel.Team, error)
}

type team struct {
	userRepo   repository.UserRepo
	authHelper helpers.Auth
}

func (svc *team) My(ctx context.Context) (*graphmodel.Team, error) {
	logger.Log.Info("verifying auth for the user")
	userID, isAuthed := svc.authHelper.IsAuthorized(ctx, 0)
	if !isAuthed {
		return nil, apiutils.HandleError(ctx, constants.Unauthorized, errors.New(constants.Unauthorized))
	}

	// just fetch the team with the user here since it will anyways needed to get the team id
	logger.Log.Info("getting user by id")
	user, err := svc.userRepo.FindOne(ctx, models.UserQuery{
		User: models.User{
			Base: models.Base{
				ID: userID,
			},
		},
	}, true)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, apiutils.HandleError(ctx, constants.NotFound, errors.New(constants.UserNotFound))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	team := user.Team.Serialize()
	team.User = user.Serialize()

	return team, nil
}

func NewTeam(
	userRepo repository.UserRepo,
	authHelper helpers.Auth,
) Team {
	return &team{
		userRepo:   userRepo,
		authHelper: authHelper,
	}
}
