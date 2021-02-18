package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Zzocker/bookolab/api"
	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/core"
	"github.com/Zzocker/bookolab/middleware"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/gin-gonic/gin"
)

// CreateAndRun : creates and run the server
func CreateAndRun(lg blog.Logger, conf *config.ApplicationConf) {
	lg.Infof("creating and running the server")
	// builder all cores
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := core.Build(ctx, lg, *conf)
	if err != nil {
		lg.Errorf("error while building the cores : %v", err)
		os.Exit(1)
	}
	lg.Infof("successfully built all the cores")
	///////////////////////////////////////////
	engin := gin.New()
	// engin.Use(gin.Logger())
	engin.Use(gin.Recovery())

	public := engin.Group("/v1/public")
	private := engin.Group("/v1/api")
	private.Use(middleware.OAuth())
	api.RegisterHandlers(lg, public, private)
	start(lg, engin, conf.Port)
}

func start(lg blog.Logger, engin *gin.Engine, port int) {
	lg.Infof("server started on port=%d", port)
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      engin,
		WriteTimeout: 5 * time.Second,
	}
	srv.ListenAndServe()
}
