package ports

import (
	"context"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/errors"
)

// UserDatastore : port through which userCore will interect with user database
type UserDatastore interface {
	Store(ctx context.Context, user model.User) errors.E
	Get(ctx context.Context, username string) (*model.User, errors.E)
	Update(ctx context.Context, user model.User) errors.E
	Delete(ctx context.Context, username string) errors.E
	Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber int64) ([]model.User, errors.E)
}

// ImageStore : port though which imageCore will interect with image database
type ImageStore interface {
	Store(ctx context.Context, img model.Image) errors.E
	Get(ctx context.Context, imageID string) (*model.Image, errors.E)
	Update(ctx context.Context, img model.Image) errors.E
	Delete(ctx context.Context, imgID string) errors.E
	Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber int64) ([]model.Image, errors.E)
}
