package core

// CreateCommentInput represents input arg for create a new comment
type CreateCommentInput struct {
	// Comment is actual data of comment
	Comment string `json:"comment"`

	// CommentOn : identifier of the entity on which comment is to be made
	// like username ,bookid , comment_id
	CommentOn string `json:"comment_on"`
}

// UpdateCommentInput represents input arg for updating the comment comment
type UpdateCommentInput struct {
	UpdatedComment string `json:"updated_comment"`
}
