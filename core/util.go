package core

import (
	"crypto/md5"
	"encoding/hex"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func hash(s string) string {
	return hex.EncodeToString(md5.New().Sum([]byte(s)))
}

func id() string {
	return primitive.NewObjectID().Hex()
}
