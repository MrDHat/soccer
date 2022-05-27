package middlewares

import (
	"errors"
	"soccer-manager/constants"
	"soccer-manager/jwt"
	"soccer-manager/logger"
	"soccer-manager/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Auth is the middleware for authentication
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader != "" {
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) == 2 {
				ctx := c.Request.Context()
				token := tokenParts[1]
				if len(token) == 0 {
					logger.Log.WithError(errors.New(constants.Unauthorized)).Error(constants.Unauthorized)
					return
				} else {
					isAuthorized := true
					jwtInfo := &jwt.JwtKey{TokenString: token}
					err := jwtInfo.ParseJWT()
					if err != nil {
						isAuthorized = false
						logger.Log.WithError(err).Error(constants.Unauthorized)
					}

					userID := int64(jwtInfo.Claims["id"].(float64))
					expirationDate, err := time.Parse(time.RFC3339, jwtInfo.Claims["exp"].(string))
					if err != nil {
						isAuthorized = false
						logger.Log.WithError(err).Error(constants.Unauthorized)
					}
					if expirationDate.Before(time.Now()) {
						isAuthorized = false
						logger.Log.WithError(err).Error(constants.Unauthorized)
					}
					if isAuthorized {
						ctx = utils.WithUserID(ctx, userID)
						ctx = utils.WithUserToken(ctx, token)
					}
				}

				c.Request = c.Request.WithContext(ctx)
				c.Next()
			}
		}
	}
}
