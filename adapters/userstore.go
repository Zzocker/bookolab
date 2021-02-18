package adapters

import (
	"context"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/datastore"
	"github.com/Zzocker/bookolab/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)



type userStore struct {
	ds datastore.SmartDS
}

func (u *userStore) Store(ctx context.Context, user model.User) errors.E {
	return u.ds.Store(ctx, user)
}
func (u *userStore) Get(ctx context.Context, username string) (*model.User, errors.E) {
	raw, err := u.ds.Get(ctx, map[string]interface{}{
		"_id": username,
	})
	if err != nil {
		return nil, err
	}
	var user model.User
	bson.Unmarshal(raw, &user)
	return &user, nil
}
func (u *userStore) Update(ctx context.Context, user model.User) errors.E {
	return u.ds.Update(ctx, map[string]interface{}{
		"_id": user.Username,
	}, user)
}
func (u *userStore) Delete(ctx context.Context, username string) errors.E {
	return u.ds.Delete(ctx, map[string]interface{}{
		"_id": username,
	})
}
func (u *userStore) Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber int64) ([]model.User, errors.E) {
	raws, err := u.ds.Query(ctx, sortKey, query, pageNumber, defaultDocPerPage)
	if err != nil {
		return nil, err
	}
	users := make([]model.User, len(raws))
	for i := range raws {
		bson.Unmarshal(raws[i], &users[i])
	}
	return users, nil
}
