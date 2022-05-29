package api

import (
	"context"
	"errors"
	"math"

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

type Player interface {
	ListForTeam(ctx context.Context, obj *graphmodel.Team, input *graphmodel.TeamPlayerListInput) (*graphmodel.PlayerList, error)
	Update(ctx context.Context, input graphmodel.UpdatePlayerInput) (*graphmodel.Player, error)
	GetForPlayerTransfer(ctx context.Context, obj *graphmodel.PlayerTransfer) (*graphmodel.Player, error)
}

type player struct {
	playerRepo repository.PlayerRepo
	userRepo   repository.UserRepo
	authHelper helpers.Auth
}

func (svc *player) ListForTeam(ctx context.Context, obj *graphmodel.Team, input *graphmodel.TeamPlayerListInput) (*graphmodel.PlayerList, error) {
	res := &graphmodel.PlayerList{
		Data: make([]*graphmodel.Player, 0),
	}

	// no need to auth here, because the team is already authenticated
	playerQ := models.PlayerQuery{
		Player: models.Player{
			Team: &models.Team{
				Base: models.Base{
					ID: obj.ID,
				},
			},
		},
	}
	fetchCount := false
	if input != nil && input.Pagination != nil {
		if input.Pagination.Limit != 0 && input.Pagination.Page != 0 {
			playerQ.Limit = &input.Pagination.Limit
			playerQ.Page = &input.Pagination.Page

			// count needs to be fetched only if pagination is asked for
			fetchCount = true
		}
	}
	players, totalRecords, err := svc.playerRepo.FindAll(ctx, playerQ, false, fetchCount)
	if err != nil {
		if err == orm.ErrNoRows {
			return res, nil
		}
		return res, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	if input != nil && input.Pagination != nil {
		if input.Pagination.Limit != 0 && input.Pagination.Page != 0 {
			totalPage := int64(math.Ceil(float64(totalRecords) / float64(*playerQ.Limit)))
			res.CurrentPage = playerQ.Page
			res.TotalPage = &totalPage
			res.TotalRecords = &totalRecords
		}
	}

	for i := range players {
		p := players[i].Serialize()
		p.Team = obj
		res.Data = append(res.Data, p)
	}

	return res, nil
}

func (svc *player) Update(ctx context.Context, input graphmodel.UpdatePlayerInput) (*graphmodel.Player, error) {
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
	}, false)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, apiutils.HandleError(ctx, constants.NotFound, errors.New(constants.UserNotFound))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	logger.Log.Info("getting player by id")
	p, err := svc.playerRepo.FindOne(ctx, models.PlayerQuery{
		Player: models.Player{
			Base: models.Base{
				ID: input.ID,
			},
		},
	})
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, apiutils.HandleError(ctx, constants.NotFound, errors.New(constants.PlayerNotFound))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	if p.Team.ID != user.Team.ID {
		return nil, apiutils.HandleError(ctx, constants.Unauthorized, errors.New(constants.Unauthorized))
	}

	logger.Log.Info("updating player")
	p.FirstName = *utils.CheckNullAndSetString(&p.FirstName, input.FirstName)
	p.LastName = *utils.CheckNullAndSetString(&p.LastName, input.LastName)
	p.Country = *utils.CheckNullAndSetString(&p.Country, input.Country)

	err = svc.playerRepo.Update(ctx, p, []string{})
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, apiutils.HandleError(ctx, constants.NotFound, errors.New(constants.PlayerNotFound))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	return p.Serialize(), nil
}

func (svc *player) GetForPlayerTransfer(ctx context.Context, obj *graphmodel.PlayerTransfer) (*graphmodel.Player, error) {
	p, err := svc.playerRepo.FindOne(ctx, models.PlayerQuery{
		Player: models.Player{
			Base: models.Base{
				ID: obj.PlayerID,
			},
		},
	})
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, apiutils.HandleError(ctx, constants.NotFound, errors.New(constants.PlayerNotFound))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	// set it to zero for security reasons
	p.CurrentValueInDollars = 0
	return p.Serialize(), nil
}

func NewPlayer(
	playerRepo repository.PlayerRepo,
	userRepo repository.UserRepo,
	authHelper helpers.Auth,
) Player {
	return &player{
		playerRepo: playerRepo,
		userRepo:   userRepo,
		authHelper: authHelper,
	}
}
