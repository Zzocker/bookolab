package util

import (
	"context"

	"github.com/Zzocker/bookolab/pkg/blog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ctxKey uint8

const (
	CtxKeyUsername ctxKey = iota + 1
	ctxKeyRequest
)

// LoggerFromCtx will return logger with field set to request ID
func LoggerFromCtx(ctx context.Context, lg blog.Logger) blog.Logger {
	return blog.NewWithFields(lg, map[string]interface{}{
		"requestID": ctx.Value(ctxKeyRequest),
	})
}

// SetRequestID : will set new request to ctx
// request id to be used for logging
func SetRequestID(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKeyRequest, primitive.NewObjectID().Hex()) // TODO : look out for this new id
}
