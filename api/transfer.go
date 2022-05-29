package api

import (
	"context"
	"errors"

	"soccer-manager/api/helpers"
	apiutils "soccer-manager/api/utils"
	"soccer-manager/api/validators"
	"soccer-manager/constants"
	"soccer-manager/db/models"
	graphmodel "soccer-manager/graph/model"
	"soccer-manager/logger"
	"soccer-manager/repository"

	"github.com/astaxie/beego/orm"
)

type Transfer interface {
	Create(ctx context.Context, input graphmodel.CreateTransferInput) (*graphmodel.PlayerTransfer, error)
}

type transfer struct {
	validator          validators.Transfer
	userRepo           repository.UserRepo
	playerRepo         repository.PlayerRepo
	playerTransferRepo repository.PlayerTransferRepo
	authHelper         helpers.Auth
}

func (svc *transfer) Create(ctx context.Context, input graphmodel.CreateTransferInput) (*graphmodel.PlayerTransfer, error) {
	logger.Log.Info("validating input")
	if err := svc.validator.CreateInput(input); err != nil {
		return nil, apiutils.HandleError(ctx, constants.InvalidRequestData, err)
	}

	logger.Log.Info("authenticating user")
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
				ID: input.PlayerID,
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

	// TODO: verify if the player is already being transferred

	logger.Log.Info("creating player transfer")
	t := models.PlayerTransfer{
		AmountInDollars: input.AmountInDollars,
		OwnerTeam:       p.Team,
		Player:          p,
	}
	err = svc.playerTransferRepo.Create(ctx, &t)
	if err != nil {
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}
	res := t.Serialize()
	p.TransferStatus = string(constants.PlayerTransferStatusOnSale)
	res.Player = p.Serialize()

	return res, nil

}

func NewTransfer(
	validator validators.Transfer,
	userRepo repository.UserRepo,
	playerRepo repository.PlayerRepo,
	playerTransferRepo repository.PlayerTransferRepo,
	authHelper helpers.Auth,
) Transfer {
	return &transfer{
		validator:          validator,
		userRepo:           userRepo,
		playerRepo:         playerRepo,
		playerTransferRepo: playerTransferRepo,
		authHelper:         authHelper,
	}
}
