package api

import (
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/gin-gonic/gin"
)

type handlerRegister interface {
	register(lg blog.Logger, public, private *gin.RouterGroup)
}

var (
	registerFactory = []handlerRegister{
		userRouterBuilder{},
	}
)

// RegisterHandlers :
func RegisterHandlers(lg blog.Logger, public, private *gin.RouterGroup) {
	lg.Infof("Registering all http handlers")
	for i := range registerFactory {
		registerFactory[i].register(lg, public, private)
	}
	lg.Infof("Successfully registered all http handlers")
}
