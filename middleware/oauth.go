package middleware

import (
	"net/http"

	"github.com/Zzocker/bookolab/core"
	"github.com/gin-gonic/gin"
)

// OAuth is auth middleware responsible for authorizing user
func OAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := core.GetTokenCore().CheckAccessToken(c, c.GetHeader("Authorization"))
		if err != nil {
			c.Status(http.StatusUnauthorized)
		} else {
			c.Set("USERNAME", username)
			c.Next()
		}
	}
}
