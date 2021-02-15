package datastore

import (
	"context"
	"testing"
	"time"

	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ds, err := newMongoDS(ctx, blog.NewTestLogger(), conf)
	assert.NotNil(t, ds)
	assert.NoError(t, err)
}

type testStruct struct {
	ID string `bson:"_id"`
	F2 string `bson:"f2"`
}

func TestInsertWithCustomID(t *testing.T) {
	id := primitive.NewObjectID().Hex()
	v := testStruct{
		ID: id,
		F2: "name",
	}
	err := ds.Store(context.Background(), v)
	assert.NoError(t, err)

	raw, err := ds.Get(context.Background(), map[string]interface{}{
		"_id": id,
	})
	var out testStruct
	bson.Unmarshal(raw, &out)
	assert.Equal(t, v, out)
}

func TestUpdateMatching(t *testing.T) {
	err := ds.UpdateMatching(context.Background(), map[string]interface{}{
		"f2": "name",
	}, map[string]interface{}{
		"$set": map[string]interface{}{
			"f2": "new_Name",
		},
	})
	assert.NoError(t, err)
}
