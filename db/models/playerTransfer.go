package models

type PlayerTransfer struct {
	Base
	AmountInDollars int64   `json:"amount_in_dollars" orm:"column(amount_in_dollars)"`
	OwnerTeam       *Team   `json:"owner_team" orm:"rel(fk)"`
	Player          *Player `json:"player" orm:"rel(fk)"`
}
