package validators

import (
	"errors"

	"soccer-manager/constants"
	graphmodel "soccer-manager/graph/model"
	"soccer-manager/logger"
)

type Transfer interface {
	CreateInput(input graphmodel.CreateTransferInput) error
}

type transfer struct{}

func (v *transfer) CreateInput(input graphmodel.CreateTransferInput) error {
	var (
		groupError = "VALIDATE_CREATE_INPUT"
		err        error
	)

	if input.PlayerID == 0 {
		err = errors.New(constants.TransferPlayerIDEmpty)
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	if input.AmountInDollars <= 0 {
		err = errors.New(constants.TransferAmountInvalid)
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	return nil
}

func NewTransfer() Transfer {
	return &transfer{}
}
