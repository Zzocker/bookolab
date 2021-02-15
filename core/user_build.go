package core

import (
	"context"

	"github.com/Zzocker/bookolab/adapters"
	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
)

type userCoreBuilder struct{}

func (userCoreBuilder) build(ctx context.Context, lg blog.Logger, conf config.ApplicationConf) error {
	lg.Infof("Building user core")
	uStore, err := adapters.NewUserStore(ctx, lg, conf.Cores.User.UserStore)
	if err != nil {
		return err
	}
	uCore = &userCore{
		uStore: uStore,
	}
	return nil
}
