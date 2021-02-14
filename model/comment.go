package model

// CommentType represents various type of comment
// commentOnUser
// commentOnComment
type CommentType uint8

// Comment :
type Comment struct {
	ID string `json:"id" bson:"_id"`
	By string `json:"by" bson:"by"`
	On string `json:"on" bson:"on"`
}
