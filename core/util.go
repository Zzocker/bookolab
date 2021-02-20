package core

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/Zzocker/bookolab/pkg/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func hash(s string) string {
	return hex.EncodeToString(md5.New().Sum([]byte(s)))
}

func id() string {
	return primitive.NewObjectID().Hex()
}

func unWrapUsername(ctx context.Context) string {
	return ctx.Value(util.CtxKeyUsername).(string)
}
