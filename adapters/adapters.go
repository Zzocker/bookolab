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
func NewUserStore(ctx context.Context, lg blog.Logger, cfg config.DatastoreConf) (ports.UserDatastore, error) {
	ds, err := datastore.NewSmartDS(ctx, lg, cfg)
	if err != nil {
		return nil, err
	}
	return &userStore{
		ds: ds,
	}, nil
}

// NewImageStore :
func NewImageStore(ctx context.Context, lg blog.Logger, cfg config.DatastoreConf) (ports.ImageStore, error) {
	ds, err := datastore.NewSmartDS(ctx, lg, cfg)
	if err != nil {
		return nil, err
	}
	return &imageStore{
		ds: ds,
	}, nil
}

// NewCommentStore :
func NewCommentStore(ctx context.Context, lg blog.Logger, cfg config.DatastoreConf) (ports.CommentStore, error) {
	ds, err := datastore.NewSmartDS(ctx, lg, cfg)
	if err != nil {
		return nil, err
	}
	return &commentStore{
		ds: ds,
	}, nil
}

// NewTokenStore :
func NewTokenStore(ctx context.Context, lg blog.Logger, cfg config.DatastoreConf) (ports.TokenStore, error) {
	ds, err := datastore.NEwDumbDS(ctx, lg, cfg)
	if err != nil {
		return nil, err
	}
	return &tokenStore{
		ds: ds,
	}, nil
}

// NewBookStore :
func NewBookStore(ctx context.Context, lg blog.Logger, cfg config.DatastoreConf) (ports.BookStore, error) {
	ds, err := datastore.NewSmartDS(ctx, lg, cfg)
	if err != nil {
		return nil, err
	}
	return &bookStore{
		ds: ds,
	}, nil
}
