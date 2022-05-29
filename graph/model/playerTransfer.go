package graphmodel

type PlayerTransfer struct {
	ID              int64                 `json:"id"`
	PlayerID        int64                 `json:"playerId"`
	AmountInDollars *int64                `json:"amountInDollars"`
	OwnerTeamID     *int64                `json:"ownerTeamId"`
	Status          *PlayerTransferStatus `json:"status"`
}
