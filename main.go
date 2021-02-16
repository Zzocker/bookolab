package main

import (
	"flag"

	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/server"
)

var (
	configPath = flag.String("config", "config/local.yaml", "configuration path location")
)

func main() {
	flag.Parse()
	conf, err := config.Load(*configPath)
	if err != nil {
		panic(conf)
	}
	lg := blog.New(conf.Level)
	lg.Infof("configuration loaded from %s", *configPath)
	server.CreateAndRun(lg, conf)
}
