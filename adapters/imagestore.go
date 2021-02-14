package adapters

import (
	"context"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/datastore"
	"github.com/Zzocker/bookolab/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type imageStore struct {
	ds datastore.SmartDS
}

func (i *imageStore) Store(ctx context.Context, img model.Image) errors.E {
	return i.ds.Store(ctx, img)
}
func (i *imageStore) Get(ctx context.Context, imageID string) (*model.Image, errors.E) {
	raw, err := i.ds.Get(ctx, map[string]interface{}{
		"_id": imageID,
	})
	if err != nil {
		return nil, err
	}
	var image model.Image
	bson.Unmarshal(raw, &image)
	return &image, nil
}
func (i *imageStore) Update(ctx context.Context, img model.Image) errors.E {
	return i.ds.Update(ctx, map[string]interface{}{
		"_id": img.ID,
	}, img)
}
func (i *imageStore) Delete(ctx context.Context, imgID string) errors.E {
	return i.ds.Delete(ctx, map[string]interface{}{
		"_id": imgID,
	})
}
func (i *imageStore) Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber int64) ([]model.Image, errors.E) {
	raws, err := i.ds.Query(ctx, sortKey, query, pageNumber, defaultDocPerPage)
	if err != nil {
		return nil, err
	}
	images := make([]model.Image, len(raws))
	for i := range raws {
		bson.Unmarshal(raws[i], &images[i])
	}
	return images, nil
}
func (i *imageStore) DeleteAll(ctx context.Context, owner string) errors.E {
	return i.ds.DeleteMatching(ctx, map[string]interface{}{
		"owner": owner,
	})
}