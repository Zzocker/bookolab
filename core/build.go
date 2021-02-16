package core

import (
	"context"

	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
)

type coreBuilder interface {
	build(ctx context.Context, lg blog.Logger, conf config.ApplicationConf) error
}

var (
	builderFactory = []coreBuilder{
		userCoreBuilder{},
		bookCoreBuilder{},
	}
)

// Build : will all core presents in builderFactory
func Build(ctx context.Context, lg blog.Logger, conf config.ApplicationConf) error {
	lg.Infof("Building All cores")
	for i := range builderFactory {
		if err := builderFactory[i].build(ctx, lg, conf); err != nil {
			return err
		}
	}
	lg.Infof("Successfully built %d core", len(builderFactory))
	return nil
}
