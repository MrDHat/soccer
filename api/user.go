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
	"soccer-manager/utils"

	"github.com/astaxie/beego/orm"
	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/thoas/go-funk"
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
	authHelper    helpers.Auth
}

func (svc *user) generateToken(u *models.User) (string, error) {
	groupError := "GENERATE_TOKEN"

	logger.Log.Info("generating bearer token for the user")
	jwtInfo := &jwt.JwtKey{
		Claims: jwtLib.MapClaims{
			"createdAt": time.Now().Unix(),
			"exp":       time.Now().Add(time.Second * time.Duration(config.JWTExpirySeconds())),
			"id":        u.ID,
		},
	}

	err := jwtInfo.GenerateJWT()
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return "", err
	}

	return jwtInfo.TokenString, nil
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

	res.Token, err = svc.generateToken(&u)
	// if token generation fails for some reason, no need to fail the entire thing since the f/e can always ask for another token
	if err != nil {
		logger.Log.WithError(err).Error(constants.InternalServerError)
	}

	return res, nil
}
func (svc *user) Login(ctx context.Context, input graphmodel.LoginInput) (*graphmodel.LoginResponse, error) {
	res := &graphmodel.LoginResponse{}

	logger.Log.Info("validating input")
	err := svc.userValidator.LoginInput(input)
	if err != nil {
		return nil, apiutils.HandleError(ctx, constants.InvalidRequestData, err)
	}
	logger.Log.Info("input is valid")

	logger.Log.Info("finding user by email")
	user, err := svc.userRepo.FindOne(ctx, models.UserQuery{
		User: models.User{
			Email: input.Email,
		},
	}, true)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, apiutils.HandleError(ctx, constants.InvalidRequestData, errors.New(constants.UserNotFound))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	logger.Log.Info("validating user's password")
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, apiutils.HandleError(ctx, constants.Unauthorized, err)
	}

	res.Token, err = svc.generateToken(user)
	if err != nil {
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}
	res.User = user.Serialize()

	return res, nil
}
func (svc *user) Me(ctx context.Context) (*graphmodel.User, error) {
	logger.Log.Info("verifying auth for the user")
	userID, isAuthed := svc.authHelper.IsAuthorized(ctx, 0)
	if !isAuthed {
		return nil, apiutils.HandleError(ctx, constants.Unauthorized, errors.New(constants.Unauthorized))
	}

	// check if team is being asked for in the user details
	fields := utils.GraphQLPreloads(ctx)
	fetchRelatedToUser := funk.ContainsString(fields, "team")

	logger.Log.Info("getting user by id")
	user, err := svc.userRepo.FindOne(ctx, models.UserQuery{
		User: models.User{
			Base: models.Base{
				ID: userID,
			},
		},
	}, fetchRelatedToUser)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, apiutils.HandleError(ctx, constants.NotFound, errors.New(constants.UserNotFound))
		}
		return nil, apiutils.HandleError(ctx, constants.InternalServerError, err)
	}

	return user.Serialize(), nil
}

func NewUser(
	userRepo repository.UserRepo,
	userValidator validators.User,
	teamHelper helpers.Team,
	authHelper helpers.Auth,
) User {
	return &user{
		userRepo:      userRepo,
		userValidator: userValidator,
		teamHelper:    teamHelper,
		authHelper:    authHelper,
	}
}
