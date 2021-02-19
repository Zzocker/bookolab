package middleware

import (
	"net/http"

	"github.com/Zzocker/bookolab/core"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/util"
	"github.com/gin-gonic/gin"
)

// OAuth is auth middleware responsible for authorizing user
func OAuth(lg blog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := core.GetTokenCore().CheckAccessToken(c, c.GetHeader("Authorization"))
		if err != nil {
			c.Status(http.StatusUnauthorized)
		} else {
			c.Request = c.Request.WithContext(util.SetRequestID(c.Request.Context()))
			util.LoggerFromCtx(c.Request.Context(), lg).Infof("endpoint call=%s", c.Request.URL.Path)
			c.Set("USERNAME", username)
			// start := time.Now()
			c.Next()
			// lg.Debugf("")
		}
	}
}
