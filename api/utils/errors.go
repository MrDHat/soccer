package utils

import (
	"context"
	"errors"

	"soccer-manager/config"
	"soccer-manager/constants"
	"soccer-manager/instance"
	"soccer-manager/utils"

	"soccer-manager/logger"

	"github.com/astaxie/beego/orm"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// HandleError is the method for returning the user facing error
func HandleError(c context.Context, errType string, err error) error {
	req, _ := utils.RequestFromContext(c)
	logger.Log.WithField("request", req).WithError(err).Error(errType)

	if !config.IsDevEnv() && errType == constants.InternalServerError {
		notice := instance.Airbrake().Notice(err, nil, 0)
		notice.Context["body"] = req
		instance.Airbrake().Notify(notice, nil)
	}

	if errType != constants.InvalidRequestData {
		err = errors.New(constants.ErrorCode[errType])
	}
	if err == orm.ErrNoRows {
		err = errors.New(constants.ErrorCode[constants.NotFound])
	}

	errToReturn := constants.ErrorCode[err.Error()]

	if errToReturn == "" {
		errToReturn = constants.ErrorCode[errType]
	}

	graphQLErr := &gqlerror.Error{
		Extensions: map[string]interface{}{
			"code":    errToReturn,
			"message": constants.ErrorString[errType],
		},
	}

	graphQLErr.Message = errType

	return graphQLErr
}
