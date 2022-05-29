package helpers

import (
	"context"
	"strings"

	"soccer-manager/constants"
	"soccer-manager/db/models"
	"soccer-manager/utils"
)

type Team interface {
	CreateRandom(ctx context.Context) models.Team
}

type team struct{}

func (h *team) createRandomPlayer(playerType constants.PlayerType) *models.Player {
	fName := utils.RandomName()
	if len(strings.Split(fName, " ")) > 1 {
		fName = strings.Split(fName, " ")[0]
	}
	lName := utils.RandomName()
	if len(strings.Split(lName, " ")) > 1 {
		lName = strings.Split(lName, " ")[0]
	}
	p := models.Player{
		FirstName:             fName,
		LastName:              lName,
		Age:                   utils.RandomAge(),
		Country:               utils.RandomCountry(),
		CurrentValueInDollars: int64(constants.DefaultPlayerAmount),
		PlayerType:            string(playerType),
		TransferStatus:        string(constants.PlayerTransferStatusOwned),
	}

	return &p
}

func (h *team) CreateRandom(ctx context.Context) models.Team {
	// create a random team for the user
	teamCountry := utils.RandomCountry()
	t := models.Team{
		Name:                     utils.RandomName(),
		Country:                  &teamCountry,
		RemainingBudgetInDollars: constants.DefaultTeamAllotment,
	}

	t.Players = make([]*models.Player, 0)

	for pType, num := range constants.DefaultTeamPlayerMapping {
		for i := 0; i < num; i++ {
			t.Players = append(t.Players, h.createRandomPlayer(pType))
		}
	}

	return t
}

func NewTeam() Team {
	return &team{}
}
