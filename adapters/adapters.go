package adapters

import (
	"context"

	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/datastore"
	"github.com/Zzocker/bookolab/ports"
)

const (
	defaultDocPerPage = 10
)

// NewUserStore :
func NewUserStore(ctx context.Context, lg blog.Logger, cfg config.DatastoreConf) ports.UserDatastore {
	return &userStore{
		ds: datastore.NewSmartDS(ctx, lg, cfg),
	}
}

// NewImageStore :
func NewImageStore(ctx context.Context, lg blog.Logger, cfg config.DatastoreConf) ports.ImageStore {
	return &imageStore{
		ds: datastore.NewSmartDS(ctx, lg, cfg),
	}
}

// NewCommentStore :
func NewCommentStore(ctx context.Context, lg blog.Logger, cfg config.DatastoreConf) ports.CommentStore {
	return &commentStore{
		ds: datastore.NewSmartDS(ctx, lg, cfg),
	}
}
