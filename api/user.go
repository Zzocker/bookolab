package api

import (
	"net/http"
	"strconv"

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
	private.PATCH("/user/profile", uAPI.update)
	private.DELETE("/user/profile", uAPI.delete)

	private.GET("/users", uAPI.getWithName)
}

type userAPI struct {
	lg   blog.Logger
	core core.UserCore
}

func (u *userAPI) register(c *gin.Context) {
	res := newRes()
	var input core.UserRegisterInput
	jErr := c.ShouldBindJSON(&input)
	if jErr != nil {
		res.Status.Code = http.StatusBadRequest
		res.Status.Message = "invalid json request"
		res.send(c)
		return
	}
	err := u.core.Register(c.Request.Context(), input, c.GetHeader("secret"))
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.send(c)
}

func (u *userAPI) getUser(c *gin.Context) {
	user, err := u.core.GetUser(c.Request.Context(), c.Param("username"))
	res := newRes()
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.Data = user
	res.send(c)
}

func (u *userAPI) update(c *gin.Context) {
	defer c.Request.Body.Close()
	err := u.core.UpdateUser(c.Request.Context(), c.Request.Body)
	res := newRes()
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.send(c)
}

func (u *userAPI) delete(c *gin.Context) {
	err := u.core.DeleteUser(c.Request.Context())
	res := newRes()
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.send(c)
}

func (u *userAPI) getWithName(c *gin.Context) {
	pageNumber, sErr := strconv.Atoi(c.Query("page"))
	if sErr != nil {
		pageNumber = 1
	}
	users, err := u.core.GetUserWithName(c.Request.Context(), c.Query("name"), int64(pageNumber))
	res := newRes()
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.Data = users
	res.send(c)
}
