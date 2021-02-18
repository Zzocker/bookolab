package core

import (
	"context"
	"fmt"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/Zzocker/bookolab/pkg/errors"
	"github.com/Zzocker/bookolab/ports"
)

type tokenCore struct {
	tStore ports.TokenStore
}

func (t *tokenCore) CreateAccessToken(ctx context.Context, refreshToken string) (string, errors.E) {
	if refreshToken == "" {
		return "", errors.Init(fmt.Errorf("empty refresh token"), code.CodeInvalidArgument, "empty refresh token")
	}
	token, err := t.tStore.Get(ctx, refreshToken)
	if err != nil {
		return "", err
	}
	*token = model.NewAccessToken(id(), token.Username)
	err = t.tStore.Store(ctx, *token)
	if err != nil {
		return "", err
	}
	return token.ID, nil
}
func (t *tokenCore) CreateRefreshToken(ctx context.Context, username, password string) (string, errors.E) {
	if username == "" {
		return "", errors.Init(fmt.Errorf("empty username"), code.CodeInvalidArgument, "empty username")
	} else if password == "" {
		return "", errors.Init(fmt.Errorf("empty password"), code.CodeInvalidArgument, "empty password")
	}
	err := GetUserCore().CheckCred(ctx, username, password)
	if err != nil {
		return "", err
	}
	token := model.NewRefreshToken(id(), username)
	err = t.tStore.Store(ctx, token)
	if err != nil {
		return "", err
	}
	return token.ID, nil
}
func (t *tokenCore) CheckAccessToken(ctx context.Context, accessID string) (string, errors.E) {
	token, err := t.tStore.Get(ctx, accessID)
	if err != nil {
		return "", err
	}
	return token.Username, nil
}
