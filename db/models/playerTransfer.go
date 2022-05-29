package models

import (
	graphmodel "soccer-manager/graph/model"

	"github.com/astaxie/beego/orm"
)

type PlayerTransfer struct {
	Base
	AmountInDollars int64   `json:"amount_in_dollars" orm:"column(amount_in_dollars)"`
	Status          string  `json:"status" orm:"column(status)"`
	OwnerTeam       *Team   `json:"owner_team" orm:"rel(fk)"`
	Player          *Player `json:"player" orm:"rel(fk)"`
}

func (*PlayerTransfer) TableName() string {
	return "player_transfers"
}

func stringToPlayerTransferStatusEnum(val *string) *graphmodel.PlayerTransferStatus {
	if val != nil {
		res := graphmodel.PlayerTransferStatus(*val)
		return &res
	}
	return nil
}

func (m *PlayerTransfer) Serialize() *graphmodel.PlayerTransfer {
	res := &graphmodel.PlayerTransfer{
		ID:              m.ID,
		AmountInDollars: &m.AmountInDollars,
		Status:          stringToPlayerTransferStatusEnum(&m.Status),
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
	Limit     *int64
	Page      *int64
	SortOrder *string
	SortBy    *string
}

func init() {
	orm.RegisterModel(new(PlayerTransfer))
}
