package api

import (
	"context"

	"soccer-manager/api/helpers"
	apiutils "soccer-manager/api/utils"
	"soccer-manager/api/validator"
	"soccer-manager/constants"
	"soccer-manager/db/models"
	graphmodel "soccer-manager/graph/model"
	"soccer-manager/logger"
	"soccer-manager/repository"

	"golang.org/x/crypto/bcrypt"
)

type User interface {
	Signup(ctx context.Context, input graphmodel.SignupInput) (*graphmodel.LoginResponse, error)
	Login(ctx context.Context, input graphmodel.LoginInput) (*graphmodel.LoginResponse, error)
	Me(ctx context.Context) (*graphmodel.User, error)
}

type user struct {
	userValidator validator.User
	userRepo      repository.UserRepo
	teamHelper    helpers.Team
}

func (svc *user) Signup(ctx context.Context, input graphmodel.SignupInput) (*graphmodel.LoginResponse, error) {
	logger.Log.Info("validating input")
	err := svc.userValidator.SignupInput(input)
	if err != nil {
		return nil, apiutils.HandleError(ctx, constants.InvalidRequestData, err)
	}
	logger.Log.Info("input is valid")

	logger.Log.Info("generating bcrypt version of the user's password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}
	password := string(hashedPassword)
	logger.Log.Info("generated bcrypt version of the user's password")

	u := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: password,
	}

	t := svc.teamHelper.CreateRandom(ctx)

	// TODO: save team, players & user in a single transaction

	return nil, nil
}
func (svc *user) Login(ctx context.Context, input graphmodel.LoginInput) (*graphmodel.LoginResponse, error) {
	return nil, nil
}
func (svc *user) Me(ctx context.Context) (*graphmodel.User, error) {
	return nil, nil
}

func NewUser(
	userRepo repository.UserRepo,
	userValidator validator.User,
	teamHelper helpers.Team,
) User {
	return &user{
		userRepo:      userRepo,
		userValidator: userValidator,
		teamHelper:    teamHelper,
	}
}
