package service

import (
	"soccer-manager/api"
	"soccer-manager/api/helpers"
	"soccer-manager/api/validators"
	"soccer-manager/instance"
	"soccer-manager/repository"
)

// Services is the interface for enlosing all the services
type Services interface {
	User() api.User
}

type services struct {
	user api.User
}

func (s *services) User() api.User {
	return s.user
}

// Init intializes the services
func Init() Services {
	db := instance.DB()

	userRepo := repository.NewUserRepo(db)

	userValidator := validators.NewUser()

	teamHelper := helpers.NewTeam()

	return &services{
		user: api.NewUser(
			userRepo,
			userValidator,
			teamHelper,
		),
	}
}
