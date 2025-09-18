package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/numberwan0532/wanxzwork/task4/internal/service"
)

func CreateCommnetHandler(c *gin.Context) {
	err := service.InsertCommnet(c)
	CommonReturn(c, nil, err)
}

func GetCommentByPostIdHandler(c *gin.Context) {
	posts, err := service.GetCommentByPostId(c)
	CommonReturn(c, posts, err)
}
