package datastore

import (
	"context"
	"testing"
	"time"

	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/stretchr/testify/assert"
)

var (
	ds *mongoDS
)

func TestMongoType(t *testing.T) {
	var l interface{} = &mongoDS{}
	_, ok := l.(SmartDS)
	assert.True(t, ok)
}
func TestNewMongoDS(t *testing.T) {
	conf := config.DatastoreConf{
		URL:        "localhost:27017",
		Username:   "root",
		Password:   "password",
		Database:   "testDB",
		Collection: "testcol",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ds = newMongoDS(ctx, blog.NewTestLogger(), conf)
	assert.NotNil(t, ds)
}
