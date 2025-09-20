package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/numberwan0532/wanxzwork/task4/internal/service"
)

type PostHandler struct {
}

var postService service.PostService = service.PostService{}

func (postHandler *PostHandler) InsertPostHandler(c *gin.Context) {
	err := postService.InsertPost(c)
	CommonReturn(c, nil, err)
}

func (postHandler *PostHandler) GetPostByIdHandler(c *gin.Context) {
	data, err := postService.GetPostById(c)
	CommonReturn(c, data, err)
}

func (postHandler *PostHandler) GetAllPostHandler(c *gin.Context) {
	posts, err := postService.GetAllPost(c)
	CommonReturn(c, posts, err)
}

func (postHandler *PostHandler) UpdatePostHandler(c *gin.Context) {
	err := postService.UpdatePost(c)
	CommonReturn(c, nil, err)
}

func (postHandler *PostHandler) DeletePostByIdHandler(c *gin.Context) {
	err := postService.DeletePostById(c)
	CommonReturn(c, nil, err)
}
