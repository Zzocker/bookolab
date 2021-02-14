package adapters

import (
	"context"
	"encoding/json"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/datastore"
	"github.com/Zzocker/bookolab/pkg/errors"
)

type tokenStore struct {
	ds datastore.DumbDS
}

func (t *tokenStore) Store(ctx context.Context, token model.Token) errors.E {
	raw, _ := json.Marshal(token)
	return t.ds.Store(ctx, token.ID, raw, token.ExpireIn)
}
func (t *tokenStore) Get(ctx context.Context, tokenID string) (*model.Token, errors.E) {
	raw, err := t.ds.Get(ctx, tokenID)
	if err != nil {
		return nil, err
	}
	var token model.Token
	json.Unmarshal(raw, &token)
	return &token, nil
}
func (t *tokenStore) Delete(ctx context.Context, tokeID string) errors.E {
	return t.ds.Delete(ctx, tokeID)
}
