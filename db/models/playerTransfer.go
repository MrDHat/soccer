package models

import (
	graphmodel "soccer-manager/graph/model"

	"github.com/astaxie/beego/orm"
)

type PlayerTransfer struct {
	Base
	AmountInDollars int64   `json:"amount_in_dollars" orm:"column(amount_in_dollars)"`
	OwnerTeam       *Team   `json:"owner_team" orm:"rel(fk)"`
	Player          *Player `json:"player" orm:"rel(fk)"`
}

func (*PlayerTransfer) TableName() string {
	return "player_transfers"
}

func (m *PlayerTransfer) Serialize() *graphmodel.PlayerTransfer {
	res := &graphmodel.PlayerTransfer{
		ID:              m.ID,
		AmountInDollars: &m.AmountInDollars,
	}
	if m.OwnerTeam != nil {
		res.OwnerTeam = m.OwnerTeam.Serialize()
	}
	if m.Player != nil {
		res.Player = m.Player.Serialize()
	}
	return res
}

type PlayerTransferQuery struct {
	PlayerTransfer
}

func init() {
	orm.RegisterModel(new(PlayerTransfer))
}
