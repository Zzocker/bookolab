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
