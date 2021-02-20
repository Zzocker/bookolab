package middleware

import (
	"net/http"
	"time"

	"github.com/Zzocker/bookolab/core"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/util"
	"github.com/gin-gonic/gin"
)

// OAuth is auth middleware responsible for authorizing user
func OAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := core.GetTokenCore().CheckAccessToken(c, c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Request = c.Request.WithContext(wrapUsername(c.Request.Context(), username))
		c.Next()
	}
}

// Access :
func Access(lg blog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(util.SetRequestID(c.Request.Context()))
		l := util.LoggerFromCtx(c.Request.Context(), lg)
		l.Infof("request begain path=%s", c.Request.URL.Path)
		start := time.Now()
		c.Next()
		l.Infof("request end latency=%v status=%d", time.Since(start), c.Writer.Status())
	}
}
