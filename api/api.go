package api

import (
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/gin-gonic/gin"
)

type handlerRegister interface {
	register(lg blog.Logger, public, private *gin.RouterGroup) error
}

var (
	registerFactory = []handlerRegister{}
)

// RegisterHandlers : 
func RegisterHandlers(lg blog.Logger, public, private *gin.RouterGroup) error {
	lg.Infof("Registering all http handlers")
	for i := range registerFactory {
		if err := registerFactory[i].register(lg, public, private); err != nil {
			return err
		}
	}
	lg.Infof("Successfully registered all http handlers")
	return nil
}
