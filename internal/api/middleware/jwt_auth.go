package middleware

import (
	"GO-FJ/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := util.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := util.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, err)
					c.Abort()
					return
				}
				c.Set("user-id", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
	}
}
