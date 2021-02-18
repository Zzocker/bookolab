package api

import (
	"github.com/Zzocker/bookolab/core"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/gin-gonic/gin"
)

type tokenRouterBuilder struct{}

func (tokenRouterBuilder) register(lg blog.Logger, public, private *gin.RouterGroup) {
	tAPI := tokenAPI{
		core: core.GetTokenCore(),
		lg:   lg,
	}

	public.GET("/token/refresh/create", tAPI.createRefreshToken)
}

type tokenAPI struct {
	lg   blog.Logger
	core core.TokenCore
}

func (t *tokenAPI) createRefreshToken(c *gin.Context) {
	lg := blog.NewWithFields(t.lg, map[string]interface{}{
		"endpoint": "/token/refresh/create",
	})
	lg.Debugf("endpoint call")
	res := newRes()
	username := c.GetHeader("username")
	password := c.GetHeader("secret")
	lg = blog.NewWithFields(t.lg, map[string]interface{}{
		"username": username,
		"endpoint": "user/register",
	})
	tokenID, err := t.core.CreateRefreshToken(c, username, password)
	if err != nil {
		lg.Errorf(err.Error())
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.Data = tokenID
	lg.Infof("refresh token created")
	res.send(c)
}
