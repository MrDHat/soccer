package api

import (
	"context"
	"errors"
	"math"

	"soccer-manager/api/helpers"
	apiutils "soccer-manager/api/utils"
	"soccer-manager/api/validators"
	"soccer-manager/constants"
	"soccer-manager/db/models"
	graphmodel "soccer-manager/graph/model"
	"soccer-manager/logger"
	"soccer-manager/repository"
	"soccer-manager/utils"

	"github.com/astaxie/beego/orm"
)

type Transfer interface {
	Create(ctx context.Context, input graphmodel.CreateTransferInput) (*graphmodel.PlayerTransfer, error)
	List(ctx context.Context, input *graphmodel.PlayerTransferListInput) (*graphmodel.PlayerTransferList, error)
	BuyPlayer(ctx context.Context, input graphmodel.BuyPlayerInput) (*graphmodel.PlayerTransfer, error)
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
	if p.TransferStatus == string(constants.PlayerTransferStatusOnSale) {
		return nil, apiutils.HandleError(ctx, constants.InvalidRequestData, errors.New(constants.PlayerAlreadyBeingTransferred))
	}

	logger.Log.Info("creating player transfer")
	t := &models.PlayerTransfer{
		AmountInDollars: input.AmountInDollars,
		OwnerTeam:       p.Team,
		Player:          p,
		Status:          string(constants.TransferStatusPending),
	}
	err = svc.playerTransferRepo.Create(ctx, t)
	if err != nil {
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}
	res := t.Serialize()
	p.TransferStatus = string(constants.PlayerTransferStatusOnSale)

	return res, nil

}

func (svc *transfer) List(ctx context.Context, input *graphmodel.PlayerTransferListInput) (*graphmodel.PlayerTransferList, error) {
	var (
		res = &graphmodel.PlayerTransferList{
			Data: []*graphmodel.PlayerTransfer{},
		}
		fetchCount = false
		query      = models.PlayerTransferQuery{}
	)

	logger.Log.Info("authenticating user")
	userID, isAuthed := svc.authHelper.IsAuthorized(ctx, 0)
	if !isAuthed {
		return nil, apiutils.HandleError(ctx, constants.Unauthorized, errors.New(constants.Unauthorized))
	}

	if input != nil {
		if input.OnlyMine != nil && *input.OnlyMine {
			// fetching only my transfers
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
			query.OwnerTeam = user.Team
		}
		if input.Pagination != nil {
			if input.Pagination.Limit != 0 && input.Pagination.Page != 0 {
				query.Limit = &input.Pagination.Limit
				query.Page = &input.Pagination.Page

				// count needs to be fetched only if pagination is asked for
				fetchCount = true
			}
		}
		if input.Status != nil {
			query.Status = input.Status.String()
		}
	}

	logger.Log.Info("getting player transfers")
	transfers, totalRecords, err := svc.playerTransferRepo.FindAll(ctx, query, false, fetchCount)
	if err != nil {
		if err == orm.ErrNoRows {
			return res, nil
		}
	}
	if input != nil && input.Pagination != nil {
		if input.Pagination.Limit != 0 && input.Pagination.Page != 0 {
			totalPage := int64(math.Ceil(float64(totalRecords) / float64(*query.Limit)))
			res.CurrentPage = query.Page
			res.TotalPage = &totalPage
			res.TotalRecords = &totalRecords
		}
	}

	for i := range transfers {
		res.Data = append(res.Data, transfers[i].Serialize())
	}

	return res, nil
}

func (svc *transfer) BuyPlayer(ctx context.Context, input graphmodel.BuyPlayerInput) (*graphmodel.PlayerTransfer, error) {
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
	}, true)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, apiutils.HandleError(ctx, constants.NotFound, errors.New(constants.UserNotFound))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	logger.Log.Info("getting player transfer by id")
	t, err := svc.playerTransferRepo.FindOne(ctx, models.PlayerTransferQuery{
		PlayerTransfer: models.PlayerTransfer{
			Base: models.Base{
				ID: input.PlayerTransferID,
			},
		},
	}, true)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, apiutils.HandleError(ctx, constants.NotFound, errors.New(constants.PlayerTransferNotFound))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	if t.OwnerTeam.ID == user.Team.ID {
		return nil, apiutils.HandleError(ctx, constants.InvalidRequestData, errors.New(constants.PlayerTransferOwnerTeamError))
	}

	if t.Status != string(constants.TransferStatusPending) {
		return nil, apiutils.HandleError(ctx, constants.InvalidRequestData, errors.New(constants.PlayerTransferAlreadyComplete))
	}

	if user.Team.RemainingBudgetInDollars < t.AmountInDollars {
		return nil, apiutils.HandleError(ctx, constants.InvalidRequestData, errors.New(constants.PlayerTransferBudgetError))
	}

	logger.Log.Info("initiating player transfer")
	playervalIncreasePercent := utils.RandomValuePercentage()
	newPlayerValue := t.Player.CurrentValueInDollars + (t.Player.CurrentValueInDollars*playervalIncreasePercent)/100
	err = svc.playerTransferRepo.CompleteTransfer(ctx, t, user.Team.ID, newPlayerValue)
	if err != nil {
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}
	t.Status = string(constants.TransferStatusComplete)

	return t.Serialize(), nil
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
