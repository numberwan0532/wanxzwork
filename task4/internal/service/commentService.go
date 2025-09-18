package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/numberwan0532/wanxzwork/task4/internal/model"
)

func InsertCommnet(c *gin.Context) error {
	var comment model.Commnet
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		return err
	}
	comment.UserID = c.MustGet("userID").(uint)
	if _, err := model.GetPostById(fmt.Sprintf("%d", comment.PostID)); err != nil {
		return errors.New("文章不存在")
	}
	return model.CreateCommnet(comment)
}

func GetCommentByPostId(c *gin.Context) ([]model.Commnet, error) {
	id := c.Param("id")
	commonts, err := model.GetCommentByPostId(id)
	if err != nil {
		return commonts, err
	}
	return commonts, nil
}
