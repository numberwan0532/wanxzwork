package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/numberwan0532/wanxzwork/task4/internal/model"
)

func InsertPost(c *gin.Context) error {
	var post model.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		return err
	}
	fmt.Println("post:", post)
	post.UserID = c.MustGet("userID").(uint)
	fmt.Println("post1:", post)
	return model.CreatePost(post)
}

func GetPostById(c *gin.Context) (model.Post, error) {
	id := c.Param("id")
	return model.GetPostById(id)
}

func GetAllPost(c *gin.Context) ([]model.Post, error) {
	return model.GetAllPost(), nil
}

func UpdatePost(c *gin.Context) error {

	var post model.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		return err
	}
	post.UserID = c.MustGet("userID").(uint)
	data, err := model.GetPostById(fmt.Sprintf("%d", post.ID))
	if err != nil {
		return err
	}
	if data.UserID != post.UserID {
		return errors.New("不允许修改别人的文章")
	}
	data.Content = post.Content
	data.Title = post.Title
	return model.UpdatePost(data)
}

func DeletePostById(c *gin.Context) error {
	id := c.Param("id")
	data, err := model.GetPostById(id)
	if err != nil {
		return err
	}
	if data.UserID != c.MustGet("userID").(uint) {
		return errors.New("不允许删除别人的文章")
	}
	return model.DeletePostById(id)
}
