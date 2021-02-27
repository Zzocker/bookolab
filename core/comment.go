package core

import (
	"context"
	"fmt"
	"time"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/Zzocker/bookolab/pkg/errors"
	"github.com/Zzocker/bookolab/pkg/util"
	"github.com/Zzocker/bookolab/ports"
)

type commentCore struct {
	cStore ports.CommentStore
	lg     blog.Logger
}

func (c *commentCore) CommentOnUser(ctx context.Context, cmt CreateCommentInput) errors.E {
	lg := util.LoggerFromCtx(ctx, c.lg)
	lg.Debugf("validating comment request")
	if cmt.Comment == "" {
		lg.Errorf("invalid comment request : empty comment")
		return errors.Init(fmt.Errorf("cannot create a empty comment"), code.CodeInvalidArgument, "cannot create a empty comment")
	}
	_, err := GetUserCore().GetUser(ctx, cmt.CommentOn)
	if err != nil {
		lg.Errorf("invalid comment request : user doesn't exist to make a comment")
		return errors.Init(fmt.Errorf("comment on non-existing user not allowed"), code.CodeInvalidArgument, "comment on non-existing user not allowed")
	}
	lg.Debugf("storeing the comment")
	comment := model.NewComment(unWrapUsername(ctx), cmt.CommentOn, model.CommentTypeOnUser, cmt.Comment)
	err = c.create(ctx, comment)
	if err != nil {
		lg.Errorf("failed to stored the comment %v", err.Error())
		return err
	}
	return nil
}
func (c *commentCore) CommentOnBook(ctx context.Context, cmt CreateCommentInput) errors.E {
	return nil
}
func (c *commentCore) CommentOnComment(ctx context.Context, cmt CreateCommentInput) errors.E {
	return nil
}
func (c *commentCore) GetComment(ctx context.Context, cmtID string) (*model.Comment, errors.E) {
	return c.cStore.Get(ctx, cmtID)
}
func (c *commentCore) UpdateComment(ctx context.Context, cmtID string, updateCmt UpdateCommentInput) errors.E {
	lg := util.LoggerFromCtx(ctx, c.lg)
	lg.Debugf("validating update comment request")
	if updateCmt.UpdatedComment == "" {
		lg.Errorf("failed to validate comment update request : empty new comment")
		return errors.Init(fmt.Errorf("empty comment"), code.CodeInvalidArgument, "empty comment")
	}
	lg.Debugf("getting older version comment")
	comment, err := c.cStore.Get(ctx, cmtID)
	if err != nil {
		lg.Errorf("failed to get older comment %v", err.Error())
		return err
	}
	if comment.CommentMadeBy != unWrapUsername(ctx) {
		lg.Errorf("unauthorized to update comment made other user")
		return errors.Init(fmt.Errorf("unauthorized to update comment made other user"), code.CodeUnauthorized, "unauthorized to update comment made other user")
	}
	comment.Value = updateCmt.UpdatedComment
	comment.CreatedOn = time.Now().Unix()
	lg.Debugf("storing updated comment")
	err = c.cStore.Update(ctx, *comment)
	if err != nil {
		lg.Errorf("failed to store updated comment %v", err.Error())
		return err
	}
	return nil
}
func (c *commentCore) DeleteComment(ctx context.Context, cmtID string) errors.E {
	lg := util.LoggerFromCtx(ctx, c.lg)
	lg.Debugf("checking if this user can delete comment")
	comment, err := c.cStore.Get(ctx, cmtID)
	if err != nil {
		lg.Errorf("failed to get comment %v", err.Error())
		return err
	}
	if comment.CommentMadeBy != unWrapUsername(ctx) {
		lg.Errorf("unauthorized to delete this comment")
		return errors.Init(fmt.Errorf("unauthorized to delete"), code.CodeUnauthorized, "unauthorized to delete")
	}
	lg.Debugf("delting the comment")
	err = c.cStore.Delete(ctx, cmtID)
	if err != nil {
		lg.Errorf("failed to delete comment %v", err.Error())
		return err
	}
	return nil
}
func (c *commentCore) GetUserComment(ctx context.Context, username string) ([]model.Comment, errors.E) {
	return nil, nil
}
func (c *commentCore) GetBookComment(ctx context.Context, bookID string) ([]model.Comment, errors.E) {
	return nil, nil
}
func (c *commentCore) GetCommentComment(ctx context.Context, bookID string) ([]model.Comment, errors.E) {
	return nil, nil
}

func (c *commentCore) create(ctx context.Context, comment model.Comment) errors.E {
	return c.cStore.Store(ctx, comment)
}
func (c *commentCore) update(ctx context.Context, commentType model.CommentType, identifer string, comment model.Comment) errors.E {
	return nil
}
func (c *commentCore) delete(ctx context.Context, commentType model.CommentType, identifer string) errors.E {
	return nil
}

func (c *commentCore) deleteAll(ctx context.Context, identifer string) errors.E {
	return nil
}
