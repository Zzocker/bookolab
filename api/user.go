package api

import (
	"github.com/Zzocker/bookolab/core"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/gin-gonic/gin"
)

type userRouterBuilder struct{}

func (userRouterBuilder) register(lg blog.Logger, public, private *gin.RouterGroup) {
	uAPI := userAPI{
		lg:   lg,
		core: core.GetUserCore(),
	}
	public.POST("/user/register", uAPI.register)
}

type userAPI struct {
	lg   blog.Logger
	core core.UserCore
}

func (u *userAPI) register(c *gin.Context) {
	res := newRes()
	defer res.send(c)
}
