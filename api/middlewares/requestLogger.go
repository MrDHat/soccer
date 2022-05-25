package middlewares

import (
	"soccer-manager/utils"

	"soccer-manager/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLogger() func(c *gin.Context) {
	return func(c *gin.Context) {
		req, _ := utils.RequestFromContext(c.Request.Context())

		logger.Log.WithFields(logrus.Fields{
			"request": req,
		}).Info("request info")
		c.Next()
	}
}
