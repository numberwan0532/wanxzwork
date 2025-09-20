package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/numberwan0532/wanxzwork/task4/internal/model"
)

type PostService struct {
}

var p model.Post = model.Post{}

func (postService *PostService) InsertPost(c *gin.Context) error {
	var post model.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		return err
	}
	fmt.Println("post:", post)
	post.UserID = c.MustGet("userID").(uint)
	fmt.Println("post1:", post)
	return p.CreatePost(post)
}

func (postService *PostService) GetPostById(c *gin.Context) (model.Post, error) {
	id := c.Param("id")
	return p.GetPostById(id)
}

func (postService *PostService) GetAllPost(c *gin.Context) ([]model.Post, error) {
	return p.GetAllPost(), nil
}

func (postService *PostService) UpdatePost(c *gin.Context) error {

	var post model.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		return err
	}
	post.UserID = c.MustGet("userID").(uint)
	data, err := p.GetPostById(fmt.Sprintf("%d", post.ID))
	if err != nil {
		return err
	}
	if data.UserID != post.UserID {
		return errors.New("不允许修改别人的文章")
	}
	data.Content = post.Content
	data.Title = post.Title
	return p.UpdatePost(data)
}

func (postService *PostService) DeletePostById(c *gin.Context) error {
	id := c.Param("id")
	data, err := p.GetPostById(id)
	if err != nil {
		return err
	}
	if data.UserID != c.MustGet("userID").(uint) {
		return errors.New("不允许删除别人的文章")
	}
	return p.DeletePostById(id)
}
