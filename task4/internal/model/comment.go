package model

import (
	"github.com/numberwan0532/wanxzwork/task4/internal/dao"
	"gorm.io/gorm"
)

type Commnet struct {
	gorm.Model
	Content string `gorm:"not null" json:"content" binding:"required"`
	UserID  uint
	User    User
	PostID  uint `json:"postId" binding:"required"`
	Post    Post
}

func (c *Commnet) CreateCommnet(commnet Commnet) error {
	return dao.DB.Create(&commnet).Error
}

func (c *Commnet) GetCommentByPostId(id string) ([]Commnet, error) {
	var commonts []Commnet
	if err := dao.DB.Debug().Find(&commonts, "post_id = ?", id).Error; err != nil {
		return commonts, err
	}
	return commonts, nil
}
