package models

import (
	"github.com/astaxie/beego/orm"

	graphmodel "soccer-manager/graph/model"
)

type Player struct {
	Base
	FirstName             string `json:"first_name" orm:"column(first_name)"`
	LastName              string `json:"last_name" orm:"column(last_name)"`
	Age                   int64  `json:"age" orm:"column(age)"`
	CurrentValueInDollars int64  `json:"current_value_in_dollars" orm:"column(current_value_in_dollars)"`
	PlayerType            string `json:"player_type" orm:"column(player_type)"`
	TransferStatus        string `json:"transfer_status" orm:"column(transfer_status);default(owned)"`
	Country               string `json:"country" orm:"column(country)"`
	Team                  *Team  `json:"team" orm:"rel(fk)"`
}

func (*Player) TableName() string {
	return "players"
}

func stringToPlayerTypeEnum(val *string) *graphmodel.PlayerType {
	if val != nil {
		res := graphmodel.PlayerType(*val)
		return &res
	}
	return nil
}

func stringToTransferStatusEnum(val *string) *graphmodel.TransferStatus {
	if val != nil {
		res := graphmodel.TransferStatus(*val)
		return &res
	}
	return nil
}

func (m *Player) Serialize() *graphmodel.Player {
	res := &graphmodel.Player{
		ID:                    m.ID,
		CreatedAt:             m.CreatedAt,
		UpdatedAt:             m.UpdatedAt,
		FirstName:             &m.FirstName,
		LastName:              &m.LastName,
		Age:                   &m.Age,
		CurrentValueInDollars: &m.CurrentValueInDollars,
		Type:                  stringToPlayerTypeEnum(&m.PlayerType),
		TransferStatus:        stringToTransferStatusEnum(&m.TransferStatus),
		Country:               &m.Country,
	}
	if m.Team != nil {
		res.Team = m.Team.Serialize()
	}
	return res
}

type PlayerQuery struct {
	Player
	Limit     *int64
	Page      *int64
	SortOrder *string
	SortBy    *string
}

func init() {
	orm.RegisterModel(new(Player))
}
