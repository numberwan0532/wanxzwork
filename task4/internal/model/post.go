package model

import (
	"github.com/numberwan0532/wanxzwork/task4/internal/dao"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}

func (p *Post) CreatePost(post Post) error {
	return dao.DB.Create(&post).Error
}

func (p *Post) GetPostById(id string) (Post, error) {
	var post Post
	if err := dao.DB.Debug().First(&post, "id = ?", id).Error; err != nil {
		return post, err
	}
	return post, nil
}
func (p *Post) GetAllPost() []Post {
	var posts []Post
	dao.DB.Find(&posts)
	return posts
}

func (p *Post) UpdatePost(post Post) error {
	return dao.DB.Updates(post).Error
}

func (p *Post) DeletePostById(id string) error {
	return dao.DB.Delete(&Post{}, "id=?", id).Error
}
