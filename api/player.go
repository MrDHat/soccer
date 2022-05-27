package api

import (
	"context"
	"math"

	apiutils "soccer-manager/api/utils"
	"soccer-manager/constants"
	"soccer-manager/db/models"
	graphmodel "soccer-manager/graph/model"
	"soccer-manager/repository"

	"github.com/astaxie/beego/orm"
)

type Player interface {
	ListForTeam(ctx context.Context, obj *graphmodel.Team, input *graphmodel.TeamPlayerListInput) (*graphmodel.PlayerList, error)
}

type player struct {
	playerRepo repository.PlayerRepo
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

func NewPlayer(
	playerRepo repository.PlayerRepo,
) Player {
	return &player{
		playerRepo: playerRepo,
	}
}
