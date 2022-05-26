package validators

import (
	"errors"

	"soccer-manager/constants"
	graphmodel "soccer-manager/graph/model"
	"soccer-manager/logger"
)

type User interface {
	SignupInput(input graphmodel.SignupInput) error
}

type user struct {
}

func (v *user) SignupInput(input graphmodel.SignupInput) error {
	var (
		groupError = "VALIDATE_SIGNUP_INPUT"
		err        error
	)

	if input.Name == "" {
		err = errors.New(constants.SignupInputNameEmpty)
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	if input.Email == "" {
		err = errors.New(constants.SignupInputEmailEmpty)
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	if input.Password == "" {
		err = errors.New(constants.SignupInputPasswordEmpty)
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	if len(input.Password) < constants.MinPasswordLength {
		err = errors.New(constants.SignupInputPasswordTooShort)
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	return nil
}

func NewUser() User {
	return &user{}
}
