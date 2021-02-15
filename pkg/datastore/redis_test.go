package datastore

import (
	"context"
	"testing"
	"time"

	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/stretchr/testify/assert"
)

func TestRedisType(t *testing.T) {
	var l interface{} = &redisDS{}
	_, ok := l.(DumbDS)
	assert.True(t, ok)
}
func TestNewRedis(t *testing.T) {
	conf := config.DatastoreConf{
		URL:      "localhost:6379",
		Database: "5",
	}
	ds, err := newRedisDS(context.TODO(), blog.NewTestLogger(), conf)
	assert.NotNil(t, ds)
	assert.NoError(t, err)
}

func TestRedisStore(t *testing.T) {
	conf := config.DatastoreConf{
		URL:      "localhost:6379",
		Database: "5",
	}
	ds, _ := newRedisDS(context.TODO(), blog.NewTestLogger(), conf)
	assert.NotNil(t, ds)
	key := "redisStore"
	err := ds.Store(context.TODO(), key, []byte("pritam"), int64(10*time.Second.Seconds()))
	assert.NoError(t, err)
}

func TestRedisSStore(t *testing.T) {
	conf := config.DatastoreConf{
		URL:      "localhost:6379",
		Database: "5",
	}
	ds, _ := newRedisDS(context.TODO(), blog.NewTestLogger(), conf)
	assert.NotNil(t, ds)
	key := "redisSStore"
	err := ds.SStore(context.TODO(), key, "token1")
	assert.NoError(t, err)
	err = ds.SStore(context.TODO(), key, "token2")
	assert.NoError(t, err)
	err = ds.SStore(context.TODO(), key, "token2")
	assert.NoError(t, err)
}

func TestRedisGet(t *testing.T) {
	conf := config.DatastoreConf{
		URL:      "localhost:6379",
		Database: "5",
	}
	ds, _ := newRedisDS(context.TODO(), blog.NewTestLogger(), conf)
	assert.NotNil(t, ds)
	key := "redisGetTest"
	value := "value"
	ds.Store(context.TODO(), key, []byte(value), -1)

	out, err := ds.Get(context.TODO(), key)
	assert.NoError(t, err)
	assert.Equal(t, value, string(out))
}

func TestRedisSGet(t *testing.T) {
	conf := config.DatastoreConf{
		URL:      "localhost:6379",
		Database: "5",
	}
	ds, _ := newRedisDS(context.TODO(), blog.NewTestLogger(), conf)
	assert.NotNil(t, ds)
	key := "redisSGet"
	values := []string{"token1", "token2", "token3"}
	m := make(map[string]int, len(values))
	for _, v := range values {
		ds.SStore(context.TODO(), key, v)
		m[v] = 1
	}
	out, err := ds.SGet(context.TODO(), key)
	assert.NoError(t, err)
	for _, v := range out {
		m[v]--
	}

	for _, v := range m {
		assert.Zero(t, v)
	}
}

func TestRedisDelete(t *testing.T) {
	conf := config.DatastoreConf{
		URL:      "localhost:6379",
		Database: "5",
	}
	ds, _ := newRedisDS(context.TODO(), blog.NewTestLogger(), conf)
	assert.NotNil(t, ds)
	key := "redisDelete"
	ds.Store(context.TODO(), key, []byte("any"), -1)
	err := ds.Delete(context.TODO(), key)
	assert.NoError(t, err)
	_, err = ds.Get(context.Background(), key)
	assert.Error(t, err)
	assert.Equal(t, code.CodeNotFound, err.GetStatus())
}
