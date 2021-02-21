package util

import "go.mongodb.org/mongo-driver/bson/primitive"

func ID() string {
	return primitive.NewObjectID().Hex()
}
