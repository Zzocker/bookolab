package model

import (
	"time"

	"github.com/Zzocker/bookolab/pkg/util"
)

// CommentType represents various type of comment
// commentOnUser
// commentOnComment
type CommentType uint8

const (
	CommentTypeOnUser = iota + 1
)

// Comment :
type Comment struct {
	ID        string      `json:"id" bson:"_id"`
	By        string      `json:"by" bson:"by"`
	On        string      `json:"on" bson:"on"`
	CreatedOn int64       `json:"create_at" bson:"create_at"`
	Type      CommentType `json:"-" bson:"type"`
	Value     string      `json:"value" bson:"value"`
}

func NewComment(by, on string, cmtType CommentType, value string) Comment {
	return Comment{
		ID:        util.ID(),
		By:        by,
		On:        on,
		CreatedOn: time.Now().Unix(),
		Type:      cmtType,
		Value:     value,
	}
}
