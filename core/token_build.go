package core

import (
	"context"

	"github.com/Zzocker/bookolab/adapters"
	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
)

type tokenCoreBuilder struct{}

func (tokenCoreBuilder) build(ctx context.Context, lg blog.Logger, conf config.ApplicationConf) error {
	lg.Infof("building token core")
	tStore, err := adapters.NewTokenStore(ctx, lg, conf.Cores.Token.TokenStore)
	if err != nil {
		return err
	}
	tCore = &tokenCore{
		tStore: tStore,
		lg:     lg,
	}
	return nil
}
