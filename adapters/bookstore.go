package adapters

import (
	"context"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/datastore"
	"github.com/Zzocker/bookolab/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type bookStore struct {
	ds datastore.SmartDS
}

func (b *bookStore) Store(ctx context.Context, book model.Book) errors.E {
	return b.ds.Store(ctx, book)
}
func (b *bookStore) Get(ctx context.Context, bookID string) (*model.Book, errors.E) {
	raw, err := b.ds.Get(ctx, map[string]interface{}{
		"_id": bookID,
	})
	if err != nil {
		return nil, err
	}
	var book model.Book
	bson.Unmarshal(raw, &book)
	return &book, nil
}
func (b *bookStore) Update(ctx context.Context, book model.Book) errors.E {
	return b.ds.Update(ctx, map[string]interface{}{
		"_id": book.ID,
	}, book)
}
func (b *bookStore) Delete(ctx context.Context, bookID string) errors.E {
	return b.ds.Delete(ctx, map[string]interface{}{
		"_id": bookID,
	})
}
func (b *bookStore) Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber int64) ([]model.Book, errors.E) {
	raws, err := b.ds.Query(ctx, sortKey, query, pageNumber, defaultDocPerPage)
	if err != nil {
		return nil, err
	}
	books := make([]model.Book, len(raws))
	for i := range raws {
		bson.Unmarshal(raws[i], &books[i])
	}
	return books, nil
}
func (b *bookStore) DeleteAll(ctx context.Context, username string) errors.E {
	return b.ds.DeleteMatching(ctx, map[string]interface{}{
		"ownership.owner":   username,
		"ownership.current": username,
	})
}
