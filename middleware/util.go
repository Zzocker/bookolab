package middleware

import (
	"context"

	"github.com/Zzocker/bookolab/pkg/util"
)

func wrapUsername(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, util.CtxKeyUsername, username)
}
