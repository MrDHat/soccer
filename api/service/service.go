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
	Team() api.Team
	Player() api.Player
}

type services struct {
	user   api.User
	team   api.Team
	player api.Player
}

func (s *services) User() api.User {
	return s.user
}

func (s *services) Team() api.Team {
	return s.team
}

func (s *services) Player() api.Player {
	return s.player
}

// Init intializes the services
func Init() Services {
	db := instance.DB()

	userRepo := repository.NewUserRepo(db)
	teamRepo := repository.NewTeamRepo(db)
	playerRepo := repository.NewPlayerRepo(db)

	userValidator := validators.NewUser()

	teamHelper := helpers.NewTeam()
	authHelper := helpers.NewAuth()

	return &services{
		user: api.NewUser(
			userRepo,
			userValidator,
			teamHelper,
			authHelper,
		),
		team: api.NewTeam(
			userRepo,
			teamRepo,
			authHelper,
		),
		player: api.NewPlayer(
			playerRepo,
		),
	}
}
