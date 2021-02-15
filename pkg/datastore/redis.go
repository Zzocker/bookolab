package datastore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/Zzocker/bookolab/pkg/errors"
	"github.com/gomodule/redigo/redis"
)

type redisDS struct {
	pool *redis.Pool
	lg   blog.Logger
}

func newRedisDS(ctx context.Context, lg blog.Logger, conf config.DatastoreConf) (*redisDS, error) {
	lg.Infof("connecting redis database at %s", conf.Username)
	database, err := strconv.Atoi(conf.Database)
	if err != nil {
		return nil, fmt.Errorf("database should be integer")
	}
	pool := redis.Pool{
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			return redis.DialContext(
				ctx,
				"tcp",
				conf.URL,
				redis.DialUsername(conf.Username),
				redis.DialPassword(conf.Password),
				redis.DialDatabase(database),
			)
		},
	}
	lg.Debugf("create a new connection for ping request")
	conn, err := pool.GetContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get connection from pool")
	}
	defer conn.Close()
	lg.Debugf("making a ping request")
	if _, err := conn.Do("PING"); err != nil {
		return nil, fmt.Errorf("failed to ping %s", conf.URL)
	}
	lg.Infof("successfully connected to redis")
	return &redisDS{
		pool: &pool,
		lg:   lg,
	}, nil
}
func (r *redisDS) Store(ctx context.Context, key string, value []byte, expireIn int64) errors.E {
	r.lg.Debugf("stor key=%s TTL=%d sec", key, expireIn)
	conn, err := r.pool.DialContext(ctx)
	if err != nil {
		r.lg.Errorf("failed to get connection from pool", err)
		return errors.Init(err, code.CodeInternal, "internal database error")
	}
	defer conn.Close()
	if expireIn < 0 {
		_, err = conn.Do("SET", key, value)
	} else {
		_, err = conn.Do("SET", key, value, "EX", expireIn)
	}
	if err != nil {
		r.lg.Errorf("failed to store %s", key)
		return errors.Init(err, code.CodeInternal, "internal database error")
	}
	return nil
}
func (r *redisDS) SStore(ctx context.Context, key string, value string) errors.E {
	r.lg.Debugf("set store key=%s", key)
	conn, err := r.pool.DialContext(ctx)
	if err != nil {
		r.lg.Errorf("failed to get connection from pool", err)
		return errors.Init(err, code.CodeInternal, "internal database error")
	}
	defer conn.Close()
	_, err = conn.Do("SADD", key, value)
	if err != nil {
		r.lg.Errorf("failed to set store %s", key)
		return errors.Init(err, code.CodeInternal, "internal database error")
	}
	return nil
}
func (r *redisDS) Get(ctx context.Context, key string) ([]byte, errors.E) {
	r.lg.Debugf("get key=%s", key)
	conn, err := r.pool.DialContext(ctx)
	if err != nil {
		r.lg.Errorf("failed to get connection from pool", err)
		return nil, errors.Init(err, code.CodeInternal, "internal database error")
	}
	defer conn.Close()
	raw, err := redis.Bytes(conn.Do("GET", key))
	if err == redis.ErrNil {
		r.lg.Errorf("not found :", err)
		return nil, errors.Init(err, code.CodeNotFound, "item not found")
	} else if err != nil {
		r.lg.Errorf("not found (internal) :", err)
		return nil, errors.Init(err, code.CodeInternal, "internal database error")
	}
	return raw, nil
}
func (r *redisDS) SGet(ctx context.Context, key string) ([]string, errors.E) {
	r.lg.Debugf("get set with key=%s", key)
	conn, err := r.pool.DialContext(ctx)
	if err != nil {
		r.lg.Errorf("failed to get connection from pool", err)
		return nil, errors.Init(err, code.CodeInternal, "internal database error")
	}
	defer conn.Close()
	values, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err == redis.ErrNil {
		r.lg.Errorf("set not found :", err)
		return nil, errors.Init(err, code.CodeNotFound, "set not found")
	} else if err != nil {
		r.lg.Errorf("set not found (internal) :", err)
		return nil, errors.Init(err, code.CodeInternal, "internal database error")
	}
	return values, nil
}
func (r *redisDS) Delete(ctx context.Context, key string) errors.E {
	r.lg.Debugf("get set with key=%s", key)
	conn, err := r.pool.DialContext(ctx)
	if err != nil {
		r.lg.Errorf("failed to get connection from pool", err)
		return errors.Init(err, code.CodeInternal, "internal database error")
	}
	defer conn.Close()
	reply, err := redis.Int(conn.Do("DEL", key))
	if err != nil {
		r.lg.Errorf("internal server error : %v", err)
		return errors.Init(err, code.CodeInternal, "internal database error")
	}
	if reply != 1 {
		r.lg.Errorf("item not found : %v", err)
		return errors.Init(err, code.CodeNotFound, "item not found")
	}
	return nil
}
