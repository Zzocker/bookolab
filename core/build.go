package core

import (
	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
)

type coreBuilder interface {
	build(lg blog.Logger, conf config.ApplicationConf) error
}

var (
	builderFactory = []coreBuilder{}
)

// Build : will all core presents in builderFactory
func Build(lg blog.Logger, conf config.ApplicationConf) error {
	lg.Infof("Building All cores")
	for i := range builderFactory {
		if err := builderFactory[i].build(lg, conf); err != nil {
			return err
		}
	}
	lg.Infof("Successfully built %d core", len(builderFactory))
	return nil
}
