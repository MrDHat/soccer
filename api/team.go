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
	"soccer-manager/utils"

	"github.com/astaxie/beego/orm"
)

type Team interface {
	My(ctx context.Context) (*graphmodel.Team, error)
	Update(ctx context.Context, input graphmodel.UpdateTeamInput) (*graphmodel.Team, error)
}

type team struct {
	userRepo   repository.UserRepo
	teamRepo   repository.TeamRepo
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

func (svc *team) Update(ctx context.Context, input graphmodel.UpdateTeamInput) (*graphmodel.Team, error) {
	logger.Log.Info("verifying auth for the user")
	userID, isAuthed := svc.authHelper.IsAuthorized(ctx, 0)
	if !isAuthed {
		return nil, apiutils.HandleError(ctx, constants.Unauthorized, errors.New(constants.Unauthorized))
	}

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

	logger.Log.Info("getting team by id")
	team, err := svc.teamRepo.FindOne(ctx, models.TeamQuery{
		Team: models.Team{
			Base: models.Base{
				ID: input.ID,
			},
		},
	})
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, apiutils.HandleError(ctx, constants.NotFound, errors.New(constants.TeamNotFound))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	logger.Log.Info("verifying permissions to update the team")
	if user.Team.ID != team.ID {
		return nil, apiutils.HandleError(ctx, constants.Unauthorized, errors.New(constants.Unauthorized))
	}

	logger.Log.Info("updating team")
	team.Name = *utils.CheckNullAndSetString(&team.Name, input.Name)
	team.Country = utils.CheckNullAndSetString(team.Country, input.Country)

	err = svc.teamRepo.Update(ctx, team, []string{})
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, apiutils.HandleError(ctx, constants.NotFound, errors.New(constants.TeamNotFound))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	return team.Serialize(), nil
}

func NewTeam(
	userRepo repository.UserRepo,
	teamRepo repository.TeamRepo,
	authHelper helpers.Auth,
) Team {
	return &team{
		userRepo:   userRepo,
		teamRepo:   teamRepo,
		authHelper: authHelper,
	}
}
