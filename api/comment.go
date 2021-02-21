package api

import (
	"github.com/Zzocker/bookolab/core"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/gin-gonic/gin"
)

type commentRouterBuilder struct{}

func (commentRouterBuilder) register(lg blog.Logger, public, private *gin.RouterGroup) {
	cmt := commentAPI{}

	private.POST("/comment/on/user/:username", cmt.commentOnUser)
}

type commentAPI struct{}

func (commentAPI) commentOnUser(c *gin.Context) {
	err := core.GetUserCore().Comment(c.Request.Context(), c.Param("username"), "testcomment")
	res := newRes()
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.send(c)
}
