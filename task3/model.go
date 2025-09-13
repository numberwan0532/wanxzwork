package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	PostCount uint
	Posts     []Post
}

type Post struct {
	gorm.Model
	Title    string
	Status   string
	Comments []Comment
	UserID   uint
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	if err := tx.Find(&user, p.UserID).Error; err != nil {
		fmt.Printf("查询失败，%v", err)
	}
	var count int64
	tx.Model(&Post{}).Where("user_id = ?", p.UserID).Count(&count)
	user.PostCount = uint(count)
	tx.Model(&User{}).Where("id = ?", user.ID).Update("post_count", user.PostCount)
	return
}

type Comment struct {
	gorm.Model
	Content string
	PostId  uint
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// var post Post
	// if err := tx.First(&post, c.PostId).Error; err != nil {
	// 	fmt.Printf("查询失败，%v", err)
	// }
	var count int64
	tx.Model(&Comment{}).Where("post_id=?", c.PostId).Count(&count)
	if count == 0 {
		tx.Model(&Post{}).Where("id=?", c.PostId).Update("status", "无评论")
	}
	return
}

func createAllTable() {
	GlobalDB.AutoMigrate(&User{}, &Post{}, &Comment{})
}

func insertAllData() {
	var user = User{
		Name: "xiaoming",
		Posts: []Post{
			{Title: "wenzhang1", Comments: []Comment{{Content: "henhao"}, {Content: "buhao"}}},
			{Title: "wenzhang2", Comments: []Comment{{Content: "aaaa"}, {Content: "bbbb"}}},
		},
	}
	GlobalDB.Create(&user)
}

func findUserAllPostAndComment() {
	var user User
	if err := GlobalDB.Preload("Posts.Comments").First(&user, "name=?", "xiaoming").Error; err != nil {
		fmt.Printf("查询失败，%v", err)
	}
	// GlobalDB.Preload("Posts.Comments").First(&user, "name=?", "xiaoming1")
	fmt.Println(user)
}

func findMostCommentsPost() {
	var posts Post
	GlobalDB.Model(&Post{}).
		Select("posts.*, Count(comments.id) as count").
		Joins("left join comments on comments.post_id=posts.id").
		Group("posts.id").Order("count desc").Limit(1).Scan(&posts)
	data, err := json.Marshal(posts)
	if err != nil {
		fmt.Printf("JSON序列化失败: %v\n", err)
		return
	}
	fmt.Println(string(data))
}

func createPost() {
	post := Post{Title: "new wenzhang", Comments: []Comment{}, UserID: 3}
	GlobalDB.Save(&post)
}

func deleteComment() {
	var comment Comment
	GlobalDB.First(&comment, 1)
	GlobalDB.Delete(&comment)
}
