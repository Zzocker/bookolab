package api

import (
	"net/http"

	"github.com/Zzocker/bookolab/core"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/gin-gonic/gin"
)

type commentRouterBuilder struct{}

func (commentRouterBuilder) register(lg blog.Logger, public, private *gin.RouterGroup) {
	cmt := commentAPI{
		core: core.GetCommentCore(),
	}

	private.POST("/comment/on/user", cmt.commentOnUser)
	private.POST("/comment/on/comment", cmt.commentOnComment)

	private.GET("/comment/get/:id", cmt.getAComment)
	private.POST("/comment/update/:id", cmt.updateComment)
	private.DELETE("/comment/:id", cmt.deleteComment)

	private.GET("/comments/user/:id", cmt.getAllCommentOnUser)
	private.GET("/comments/comment/:id", cmt.getAllCommentsOnComment)
}

type commentAPI struct {
	core core.CommentCore
}

func (ct commentAPI) commentOnUser(c *gin.Context) {
	var comment core.CreateCommentInput
	jErr := c.ShouldBindJSON(&comment)
	res := newRes()
	if jErr != nil {
		res.Status.Code = http.StatusBadRequest
		res.Status.Message = "invalid json request"
		res.send(c)
		return
	}
	err := ct.core.CommentOnUser(c.Request.Context(), comment)
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.send(c)
}

func (ct commentAPI) getAComment(c *gin.Context) {
	res := newRes()
	cmt, err := ct.core.GetComment(c.Request.Context(), c.Param("id"))
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.Data = cmt
	res.send(c)
}

func (ct commentAPI) updateComment(c *gin.Context) {
	res := newRes()
	var comment core.UpdateCommentInput
	jErr := c.ShouldBindJSON(&comment)
	if jErr != nil {
		res.Status.Code = http.StatusBadRequest
		res.Status.Message = "invalid json request"
		res.send(c)
		return
	}
	err := ct.core.UpdateComment(c.Request.Context(), c.Param("id"), comment)
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.send(c)
}

func (ct commentAPI) deleteComment(c *gin.Context) {
	res := newRes()
	err := ct.core.DeleteComment(c.Request.Context(), c.Param("id"))
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.send(c)
}

func (ct commentAPI) getAllCommentOnUser(c *gin.Context) {
	res := newRes()
	comments, err := ct.core.GetUserComment(c.Request.Context(), c.Param("id"))
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.Data = comments
	res.send(c)
}

func (ct commentAPI) commentOnComment(c *gin.Context) {
	var comment core.CreateCommentInput
	jErr := c.ShouldBindJSON(&comment)
	res := newRes()
	if jErr != nil {
		res.Status.Code = http.StatusBadRequest
		res.Status.Message = "invalid json request"
		res.send(c)
		return
	}
	err := ct.core.CommentOnComment(c.Request.Context(), comment)
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.send(c)
}

func (ct commentAPI) getAllCommentsOnComment(c *gin.Context) {
	res := newRes()
	comments, err := ct.core.GetCommentComment(c.Request.Context(), c.Param("id"))
	if err != nil {
		res.Status.Code = code.ToHTTP(err.GetStatus())
		res.Status.Message = err.Message()
		res.send(c)
		return
	}
	res.Data = comments
	res.send(c)
}
