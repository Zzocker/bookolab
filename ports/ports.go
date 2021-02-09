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

// CommentStore : port though which imageCore will interect with image database
type CommentStore interface {
	Store(ctx context.Context, cmt model.Comment) errors.E
	Get(ctx context.Context, cmtID string) (*model.Comment, errors.E)
	Update(ctx context.Context, cmt model.Image) errors.E
	Delete(ctx context.Context, cmtID string) errors.E
	Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber int64) ([]model.Comment, errors.E)
}

// TokenStore : port though which imageCore will interect with image database
type TokenStore interface {
	Store(ctx context.Context, token model.Token) errors.E
	Get(ctx context.Context, tokenID string) (*model.Token, errors.E)
	Delete(ctx context.Context, tokeID string) errors.E
	// DeleteAll : delete all tokens owned by a user
	DeleteAll(ctx context.Context, userID string) errors.E
}
