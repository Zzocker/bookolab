package core

import (
	"context"

	"github.com/Zzocker/bookolab/adapters"
	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
)

type commentCoreBuilder struct{}

func (commentCoreBuilder) build(ctx context.Context, lg blog.Logger, conf config.ApplicationConf) error {
	lg.Infof("building comment core")
	cStore, err := adapters.NewCommentStore(ctx, lg, conf.Cores.Comment.CommentStore)
	if err != nil {
		return err
	}
	cCore = &commentCore{
		cStore: cStore,
		lg:     lg,
	}
	return nil
}
