package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/numberwan0532/wanxzwork/task4/internal/service"
)

type CommentHandler struct {
}

var commentService service.CommentService = service.CommentService{}

func (commentHandler *CommentHandler) CreateCommnetHandler(c *gin.Context) {
	err := commentService.InsertCommnet(c)
	CommonReturn(c, nil, err)
}

func (commentHandler *CommentHandler) GetCommentByPostIdHandler(c *gin.Context) {
	posts, err := commentService.GetCommentByPostId(c)
	CommonReturn(c, posts, err)
}
