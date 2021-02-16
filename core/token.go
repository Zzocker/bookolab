package core

import (
	"context"

	"github.com/Zzocker/bookolab/pkg/errors"
	"github.com/Zzocker/bookolab/ports"
)

type tokenCore struct {
	tStore ports.TokenStore
}

func (t *tokenCore) CreateAccessToken(ctx context.Context, refreshToken string) (string, errors.E) {
	return "", nil
}
func (t *tokenCore) CreateRefreshToken(ctx context.Context, password string) (string, errors.E) {
	return "", nil
}
func (t *tokenCore) CheckAccessToken(ctx context.Context, accessID string) (string, errors.E) {
	return "", nil
}
