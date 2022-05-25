package middlewares

import (
	"strings"

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

				c.Request = c.Request.WithContext(ctx)
				c.Next()
			}
		}
	}
}
