package datastore

import (
	"context"
	"fmt"

	"github.com/Zzocker/bookolab/config"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/Zzocker/bookolab/pkg/errors"
	"github.com/Zzocker/bookolab/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDS struct {
	lg blog.Logger
	ds *mongo.Collection
}

func newMongoDS(ctx context.Context, lg blog.Logger, conf config.DatastoreConf) (*mongoDS, error) {
	lg.Infof("connecting mongo database : mongodb://_:_@%s/%s/%s", conf.URL, conf.Database, conf.Collection)
	addrs := fmt.Sprintf("mongodb://%s:%s@%s", conf.Username, conf.Password, conf.URL)
	lg.Debugf("createing new mongo client")
	client, err := mongo.NewClient(options.Client().ApplyURI(addrs))
	if err != nil {
		return nil, fmt.Errorf("failed to create new client %+v", err)
	}
	lg.Debugf("connecting newly created client")
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect %+v", err)
	}
	lg.Debugf("pinging database")
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping %+v", err)
	}
	lg.Infof("ping successfully")
	return &mongoDS{
		ds: client.Database(conf.Database).Collection(conf.Collection),
		lg: lg,
	}, nil
}

func (m *mongoDS) Store(ctx context.Context, in interface{}) errors.E {
	_, err := m.ds.InsertOne(ctx, in)
	if isDuplicate(err) {
		return errors.Init(err, code.CodeAlreadyExists, "duplicate entry")
	} else if err != nil {
		return errors.Init(err, code.CodeInternal, "internal database error")
	}
	return nil
}
func (m *mongoDS) Get(ctx context.Context, filter map[string]interface{}) ([]byte, errors.E) {
	lg := util.LoggerFromCtx(ctx, m.lg)
	lg.Debugf("getting %v from smartDS", filter)
	reply := m.ds.FindOne(ctx, filter)
	if reply.Err() == mongo.ErrNoDocuments {
		lg.Errorf("%v", reply.Err())
		return nil, errors.Init(reply.Err(), code.CodeNotFound, "internal database error")
	} else if reply.Err() != nil {
		lg.Errorf("%v", reply.Err())
		return nil, errors.Init(reply.Err(), code.CodeInternal, "internal database error")
	}
	raw, err := reply.DecodeBytes()
	if err != nil {
		lg.Errorf("fail to decode database document %v", err)
		return nil, errors.Init(reply.Err(), code.CodeInternal, "failed to decode database document")
	}
	lg.Debugf("got %v from smartDS", filter)
	return raw, nil
}
func (m *mongoDS) Update(ctx context.Context, filter map[string]interface{}, in interface{}) errors.E {
	m.lg.Debugf("update filter=%v value", filter, in)
	reply, err := m.ds.UpdateOne(ctx, filter, bson.M{"$set": in})
	if err != nil {
		m.lg.Errorf("internal error : %v", err.Error())
		return errors.Init(err, code.CodeInternal, "internal database error")
	}
	if reply.MatchedCount != 1 {
		m.lg.Errorf("non document found with filter=%+v", filter)
		return errors.Init(err, code.CodeNotFound, "document not found")

	}
	return nil
}

func (m *mongoDS) UpdateMatching(ctx context.Context, filter map[string]interface{}, in interface{}) errors.E {
	m.lg.Debugf("update Matching filter=%v value", filter, in)
	_, err := m.ds.UpdateMany(ctx, filter, in)
	if err != nil {
		m.lg.Errorf("internal error : %v", err.Error())
		return errors.Init(err, code.CodeInternal, "internal database error")
	}
	return nil
}

func (m *mongoDS) Delete(ctx context.Context, filter map[string]interface{}) errors.E {
	m.lg.Debugf("delete %v", filter)
	reply, err := m.ds.DeleteOne(ctx, filter)
	if err != nil {
		m.lg.Errorf("internal error : %v", err.Error())
		return errors.Init(err, code.CodeInternal, "internal database error")
	}
	if reply.DeletedCount != 1 {
		m.lg.Errorf("non document found with filter=%+v", filter)
		return errors.Init(err, code.CodeNotFound, "document not found")
	}
	return nil
}
func (m *mongoDS) Query(ctx context.Context, sortingKey string, query map[string]interface{}, pageNumber, perPage int64) ([][]byte, errors.E) {
	m.lg.Debugf("query mongo filter=%v pageNumber=%d perPage=%d", query, pageNumber, perPage)
	skip := (pageNumber - 1) * perPage
	if skip < 0 {
		skip = 0
	}
	opts := options.FindOptions{
		Limit: &perPage,
		Skip:  &skip,
		Sort:  bson.D{{sortingKey, 1}},
	}
	cur, err := m.ds.Find(ctx, query, &opts)
	if err != nil {
		m.lg.Errorf("query fail : %v", err)
		return nil, errors.Init(err, code.CodeInternal, "query fail on mongo")
	}
	defer cur.Close(ctx)
	raws := make([][]byte, 0, cur.RemainingBatchLength())
	for cur.Next(ctx) {
		raws = append(raws, cur.Current)
	}
	return raws, nil
}
func (m *mongoDS) DeleteMatching(ctx context.Context, filter map[string]interface{}) errors.E {
	m.lg.Debugf("delete many %v", filter)
	_, err := m.ds.DeleteMany(ctx, filter)
	if err != nil {
		m.lg.Errorf("internal error : %v", err.Error())
		return errors.Init(err, code.CodeInternal, "internal database error")
	}
	return nil
}
func (m *mongoDS) CreateIndex(ctx context.Context, key string, unique bool) errors.E {
	m.lg.Debugf("createing mongo index key=%s isUnique=%v", key, unique)
	_, err := m.ds.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{key, 1}},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})
	if err != nil {
		m.lg.Errorf("failed to create index : %v", err)
		return errors.Init(err, code.CodeInternal, "failed to create index")
	}
	return nil
}

// helper
func isDuplicate(err error) bool {
	if mErr, ok := err.(mongo.WriteException); ok {
		for _, e := range mErr.WriteErrors {
			if e.Code == 11000 {
				return true
			}
		}
	}
	return false
}
