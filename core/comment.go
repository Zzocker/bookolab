package core

import (
	"context"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/errors"
	"github.com/Zzocker/bookolab/ports"
)

type commentCore struct {
	cStore ports.CommentStore
	lg     blog.Logger
}

func (c *commentCore) Create(ctx context.Context, comment model.Comment) errors.E {
	return c.cStore.Store(ctx, comment)
}
func (c *commentCore) Get(ctx context.Context, commentType model.CommentType, identifer string, pageNumber int64) ([]model.Comment, errors.E) {
	return c.cStore.Query(ctx, "on", map[string]interface{}{
		"by":   identifer,
		"type": commentType,
	}, pageNumber)
}
func (c *commentCore) Update(ctx context.Context, commentType model.CommentType, identifer string, comment model.Comment) errors.E {
	return nil
}
func (c *commentCore) Delete(ctx context.Context, commentType model.CommentType, identifer string) errors.E {
	return nil
}

func (c *commentCore) DeleteAll(ctx context.Context, identifer string) errors.E {
	return nil
}
