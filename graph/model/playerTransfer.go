package graphmodel

type PlayerTransfer struct {
	ID              int64                 `json:"id"`
	PlayerID        int64                 `json:"playerId"`
	AmountInDollars *int64                `json:"amountInDollars"`
	OwnerTeamID     *int64                `json:"ownerTeamId"`
	Status          *PlayerTransferStatus `json:"status"`
	CreatedAt       *int64                `json:"createdAt"`
	UpdatedAt       *int64                `json:"updatedAt"`
	CompletedAt     *int64                `json:"completedAt"`
}
