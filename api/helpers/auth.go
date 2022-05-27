package helpers

import (
	"context"
	"errors"

	"soccer-manager/constants"
	"soccer-manager/logger"
	"soccer-manager/utils"
)

type Auth interface {
	IsAuthorized(ctx context.Context, userID int64) (int64, bool)
}

type auth struct {
}

func (svc *auth) IsAuthorized(ctx context.Context, userID int64) (int64, bool) {
	groupError := "AUTH_IS_AUTHORIZED"

	loggedInUserID, err := utils.UserIDFromContext(ctx)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return 0, false
	}

	if userID != 0 {
		if userID != loggedInUserID {
			logger.Log.WithError(errors.New(constants.Unauthorized)).Error(groupError)
			return 0, false
		}
	}

	if loggedInUserID == 0 {
		logger.Log.WithError(errors.New(constants.Unauthorized)).Error(groupError)
		return 0, false
	}

	return loggedInUserID, true
}

func NewAuth() Auth {
	return &auth{}
}
