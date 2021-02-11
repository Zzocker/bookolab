package datastore

import (
	"context"
	"fmt"
	"os"

	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDS struct {
	lg blog.Logger
	ds *mongo.Collection
}

func newMongoDS(ctx context.Context, lg blog.Logger, conf config.DatastoreConf) *mongoDS {
	lg.Infof("connecting mongo database : mongodb://_:_@%s/%s/%s", conf.URL, conf.Database, conf.Collection)
	addrs := fmt.Sprintf("mongodb://%s:%s@%s", conf.Username, conf.Password, conf.URL)
	lg.Debugf("createing new mongo client")
	client, err := mongo.NewClient(options.Client().ApplyURI(addrs))
	if err != nil {
		lg.Errorf("failed to create new client %+v", err)
		os.Exit(1)
	}
	lg.Debugf("connecting newly created client")
	err = client.Connect(ctx)
	if err != nil {
		lg.Errorf("failed to connect %+v", err)
		os.Exit(1)
	}
	lg.Debugf("pinging database")
	if err = client.Ping(ctx, nil); err != nil {
		lg.Errorf("failed to ping %+v", err)
		os.Exit(1)
	}
	lg.Infof("ping successfully")
	return &mongoDS{
		ds: client.Database(conf.Database).Collection(conf.Collection),
		lg: lg,
	}
}

func (m *mongoDS) Store(ctx context.Context, in interface{}) errors.E {
	return nil
}
func (m *mongoDS) Get(ctx context.Context, key, value string) ([]byte, errors.E) {
	return nil, nil
}
func (m *mongoDS) Update(ctx context.Context, key, keyValue string, in interface{}) errors.E {
	return nil
}
func (m *mongoDS) Delete(ctx context.Context, key, keyValue string) errors.E {
	return nil
}
func (m *mongoDS) Query(ctx context.Context, sortingKey string, query map[string]interface{}, pageNumber, perPage int64) ([][]byte, errors.E) {
	return nil, nil
}
func (m *mongoDS) DeleteMatching(ctx context.Context, query map[string]interface{}) errors.E {
	return nil
}
func (m *mongoDS) CreateIndex(ctx context.Context, key string, unique bool) errors.E {
	return nil
}
