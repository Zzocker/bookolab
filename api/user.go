package api

import (
	"net/http"

	"github.com/Zzocker/bookolab/core"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/gin-gonic/gin"
)

type userRouterBuilder struct{}

func (userRouterBuilder) register(lg blog.Logger, public, private *gin.RouterGroup) {
	uAPI := userAPI{
		lg:   lg,
		core: core.GetUserCore(),
	}
	public.POST("/user/register", uAPI.register)
	// now private
	private.GET("/user/:username", uAPI.getUser)
}

type userAPI struct {
	lg   blog.Logger
	core core.UserCore
}

func (u *userAPI) register(c *gin.Context) {
	lg := blog.NewWithFields(u.lg, map[string]interface{}{
		"endpoint": "user/register",
	})
	lg.Debugf("endpoint call")
	res := newRes()
	secret := c.GetHeader("secret")
	var input core.UserRegisterInput
	jErr := c.ShouldBindJSON(&input)
	if jErr != nil {
		lg.Errorf("invalid json request : %v", jErr)
		res.Status.Code = http.StatusBadRequest
		res.Status.Message = "invalid json request"
		res.send(c)
		return
	}
	lg = blog.NewWithFields(u.lg, map[string]interface{}{
		"username": input.Username,
		"endpoint": "user/register",
	})
	err := u.core.Register(c, input, secret)
	if err != nil {
		lg.Errorf("failed to register : %v", err.Error())
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	lg.Infof("success")
	res.send(c)
}

func (u *userAPI) getUser(c *gin.Context) {
	username, _ := c.Get("USERNAME")
	lg := blog.NewWithFields(u.lg, map[string]interface{}{
		"endpoint": "user/register",
		"username": username,
	})
	lg.Debugf("endpoint call")
	user, err := u.core.GetUser(c, c.Param("username"))
	res := newRes()
	if err != nil {
		lg.Errorf("failed to user %v", err)
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	lg.Infof("got user=%s", c.Param("username"))
	res.Data = user
	res.send(c)
}
