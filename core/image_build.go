package core

import (
	"context"

	"github.com/Zzocker/bookolab/adapters"
	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
)

type imageCoreBuilder struct{}

func (imageCoreBuilder) build(ctx context.Context, lg blog.Logger, conf config.ApplicationConf) error {
	lg.Infof("building image core")
	iStore, err := adapters.NewImageStore(ctx, lg, conf.Cores.Image.ImageStore)
	if err != nil {
		return err
	}
	iCore = &imageCore{
		iStore: iStore,
	}
	return nil
}
