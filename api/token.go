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
	public.GET("/token/access/create", tAPI.createAccessToken)
}

type tokenAPI struct {
	lg   blog.Logger
	core core.TokenCore
}

func (t *tokenAPI) createRefreshToken(c *gin.Context) {
	res := newRes()
	tokenID, err := t.core.CreateRefreshToken(c.Request.Context(), c.GetHeader("username"), c.GetHeader("secret"))
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.Data = tokenID
	res.send(c)
}

func (t *tokenAPI) createAccessToken(c *gin.Context) {
	res := newRes()
	tokenID, err := t.core.CreateAccessToken(c, c.GetHeader("Refresh-Token"))
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.Data = tokenID
	res.send(c)
}
