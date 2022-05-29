package constants

type PlayerTransferStatus string

type TransferStatus string

const (
	PlayerTransferStatusOwned  PlayerTransferStatus = "owned"
	PlayerTransferStatusOnSale PlayerTransferStatus = "on_sale"

	TransferStatusPending  TransferStatus = "pending"
	TransferStatusComplete TransferStatus = "complete"
)
