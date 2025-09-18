package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/numberwan0532/wanxzwork/task4/internal/service"
)

func InsertPostHandler(c *gin.Context) {
	err := service.InsertPost(c)
	CommonReturn(c, nil, err)
}

func GetPostByIdHandler(c *gin.Context) {
	data, err := service.GetPostById(c)
	CommonReturn(c, data, err)
}

func GetAllPostHandler(c *gin.Context) {
	posts, err := service.GetAllPost(c)
	CommonReturn(c, posts, err)
}

func UpdatePostHandler(c *gin.Context) {
	err := service.UpdatePost(c)
	CommonReturn(c, nil, err)
}

func DeletePostByIdHandler(c *gin.Context) {
	err := service.DeletePostById(c)
	CommonReturn(c, nil, err)
}
