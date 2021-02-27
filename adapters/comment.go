package adapters

import (
	"context"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/datastore"
	"github.com/Zzocker/bookolab/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type commentStore struct {
	ds datastore.SmartDS
}

func (c *commentStore) Store(ctx context.Context, cmt model.Comment) errors.E {
	return c.ds.Store(ctx, cmt)
}
func (c *commentStore) Get(ctx context.Context, cmtID string) (*model.Comment, errors.E) {
	raw, err := c.ds.Get(ctx, map[string]interface{}{
		"_id": cmtID,
	})
	if err != nil {
		return nil, err
	}
	var comment model.Comment
	bson.Unmarshal(raw, &comment)
	return &comment, nil
}
func (c *commentStore) Update(ctx context.Context, cmt model.Comment) errors.E {
	return c.ds.Update(ctx, map[string]interface{}{
		"_id": cmt.ID,
	}, cmt)
}
func (c *commentStore) Delete(ctx context.Context, cmtID string) errors.E {
	return c.ds.Delete(ctx, map[string]interface{}{
		"_id": cmtID,
	})
}
func (c *commentStore) Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber int64) ([]model.Comment, errors.E) {
	raws, err := c.ds.Query(ctx, sortKey, query, pageNumber, defaultDocPerPage)
	if err != nil {
		return nil, err
	}
	comments := make([]model.Comment, len(raws))
	for i := range raws {
		bson.Unmarshal(raws[i], &comments[i])
	}
	return comments, nil
}

// TODO create two separate goroutine
// made by will be changed to bhuta
func (c *commentStore) DeleteAll(ctx context.Context, username string) errors.E {
	// now ghost will make his/her comments
	c.ds.UpdateMatching(ctx, map[string]interface{}{
		"by": username,
	}, map[string]interface{}{
		"by": "bhuta", // TODO make this global maybe in config
	})

	// delete all the comments made on this user
	// to save space
	return c.ds.DeleteMatching(ctx, map[string]interface{}{
		"on": username,
	})
}
