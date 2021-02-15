// Package datastore give access to various kind of database
// currently two kind datastore is supported
// 1. smartDS  : database with query features eg: mongo
// 2. dumbDS : database with no query features eg: redis , etcd
package datastore

import (
	"context"

	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/errors"
)

// SmartDS : represnets a datastore which support query feature
// eg : mongo
type SmartDS interface {
	Store(ctx context.Context, in interface{}) errors.E
	Get(ctx context.Context, filter map[string]interface{}) ([]byte, errors.E)
	Update(ctx context.Context, filter map[string]interface{}, in interface{}) errors.E
	UpdateMatching(ctx context.Context, query map[string]interface{}, in interface{}) errors.E
	Delete(ctx context.Context, filter map[string]interface{}) errors.E
	Query(ctx context.Context, sortingKey string, query map[string]interface{}, pageNumber, perPage int64) ([][]byte, errors.E)
	DeleteMatching(ctx context.Context, query map[string]interface{}) errors.E
	CreateIndex(ctx context.Context, key string, unique bool) errors.E
}

// DumbDS : represents a datastore which doesn't support query feature
// this type datastore don't care about value of the key ,only matters
// eg redis, etcd
// if expireIn < 0 ; means key won't expire
type DumbDS interface {
	Store(ctx context.Context, key string, value []byte, expireIn int64) errors.E
	// SStore : set store, will append set key with new value
	SStore(ctx context.Context, key string, value string) errors.E
	Get(ctx context.Context, key string) ([]byte, errors.E)
	// SGet : set get will return all values stored in the store key
	SGet(ctx context.Context, key string) ([]string, errors.E)
	// Set can also be deleted using delete
	Delete(ctx context.Context, key string) errors.E
}

// NewSmartDS :
// TODO require datastore config as argument
func NewSmartDS(ctx context.Context, lg blog.Logger, conf config.DatastoreConf) (SmartDS, error) {
	return newMongoDS(ctx, lg, conf) // TODO
}

// NEwDumbDS :
// TODO require datastore config as argument
func NEwDumbDS(ctx context.Context, lg blog.Logger, conf config.DatastoreConf) (DumbDS, error) {
	return newRedisDS(ctx, lg, conf) // TODO
}
