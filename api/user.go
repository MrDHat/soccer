package api

import (
	"context"
	"errors"
	"time"

	"soccer-manager/api/helpers"
	apiutils "soccer-manager/api/utils"
	"soccer-manager/api/validators"
	"soccer-manager/config"
	"soccer-manager/constants"
	"soccer-manager/db/models"
	graphmodel "soccer-manager/graph/model"
	"soccer-manager/jwt"
	"soccer-manager/logger"
	"soccer-manager/repository"

	jwtLib "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	Signup(ctx context.Context, input graphmodel.SignupInput) (*graphmodel.LoginResponse, error)
	Login(ctx context.Context, input graphmodel.LoginInput) (*graphmodel.LoginResponse, error)
	Me(ctx context.Context) (*graphmodel.User, error)
}

type user struct {
	userValidator validators.User
	userRepo      repository.UserRepo
	teamHelper    helpers.Team
}

func (svc *user) Signup(ctx context.Context, input graphmodel.SignupInput) (*graphmodel.LoginResponse, error) {
	res := &graphmodel.LoginResponse{}

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

	logger.Log.Info("saving users to the database")
	err = svc.userRepo.SaveWithTeamAndPlayers(ctx, &u, &t)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			return nil, apiutils.HandleError(ctx, constants.InvalidRequestData, errors.New(constants.UserAlreadyExists))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	res.User = u.Serialize()

	logger.Log.Info("generating bearer token for the user")
	jwtInfo := &jwt.JwtKey{
		Claims: jwtLib.MapClaims{
			"createdAt": time.Now().Unix(),
			"exp":       time.Now().Add(time.Second * time.Duration(config.JWTExpirySeconds())),
			"id":        u.ID,
		},
	}

	// TODO: rollback if token generation fails
	err = jwtInfo.GenerateJWT()
	if err != nil {
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}
	res.Token = jwtInfo.TokenString

	return res, nil
}
func (svc *user) Login(ctx context.Context, input graphmodel.LoginInput) (*graphmodel.LoginResponse, error) {
	return nil, nil
}
func (svc *user) Me(ctx context.Context) (*graphmodel.User, error) {
	return nil, nil
}

func NewUser(
	userRepo repository.UserRepo,
	userValidator validators.User,
	teamHelper helpers.Team,
) User {
	return &user{
		userRepo:      userRepo,
		userValidator: userValidator,
		teamHelper:    teamHelper,
	}
}
