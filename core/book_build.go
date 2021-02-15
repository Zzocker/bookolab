package core

import (
	"context"

	"github.com/Zzocker/bookolab/adapters"
	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
)

type bookCoreBuilder struct{}

func (bookCoreBuilder) build(ctx context.Context, lg blog.Logger, conf config.ApplicationConf) error {
	lg.Infof("building book core")
	bStore, err := adapters.NewBookStore(ctx, lg, conf.Cores.Book.BookStore)
	if err != nil {
		return err
	}
	bCore = &bookCore{
		bStore: bStore,
	}
	return nil
}
