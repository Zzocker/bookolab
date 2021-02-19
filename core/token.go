package core

import (
	"context"
	"fmt"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/Zzocker/bookolab/pkg/errors"
	"github.com/Zzocker/bookolab/pkg/util"
	"github.com/Zzocker/bookolab/ports"
)

type tokenCore struct {
	tStore ports.TokenStore
	lg     blog.Logger
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
	lg := util.LoggerFromCtx(ctx, t.lg)
	lg.Debugf("validating create refresh token arguments")
	if username == "" {
		lg.Errorf("empty username")
		return "", errors.Init(fmt.Errorf("empty username"), code.CodeInvalidArgument, "empty username")
	} else if password == "" {
		lg.Errorf("empty password")
		return "", errors.Init(fmt.Errorf("empty password"), code.CodeInvalidArgument, "empty password")
	}
	lg.Debugf("valideting user credentials")
	err := GetUserCore().CheckCred(ctx, username, password)
	if err != nil {
		lg.Errorf("failed to validate user : %v", err.Error())
		return "", err
	}
	token := model.NewRefreshToken(id(), username)
	lg.Debugf("storing created refresh token")
	err = t.tStore.Store(ctx, token)
	if err != nil {
		lg.Errorf("failed to store the refresh token %v", err.Error())
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
