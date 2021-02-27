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
	CommentTypeOnBook
	CommentTypeOnComment
)

// Comment :
type Comment struct {
	ID            string      `json:"id" bson:"_id"`
	CommentMadeBy string      `json:"by" bson:"comment_made_by"`
	CommentMadeOn string      `json:"on" bson:"comment_made_on"`
	CreatedOn     int64       `json:"create_at" bson:"create_at"`
	Type          CommentType `json:"-" bson:"type"`
	Value         string      `json:"value" bson:"value"`
}

func NewComment(by, on string, cmtType CommentType, value string) Comment {
	return Comment{
		ID:            util.ID(),
		CommentMadeBy: by,
		CommentMadeOn: on,
		CreatedOn:     time.Now().Unix(),
		Type:          cmtType,
		Value:         value,
	}
}
