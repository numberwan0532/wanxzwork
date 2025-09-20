package model

import (
	"github.com/numberwan0532/wanxzwork/task4/internal/dao"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

func (u *User) GetUserByUsername(username string) (User, error) {
	var user User
	if err := dao.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
func (u *User) GetUserByEmail(email string) error {
	var user User
	return dao.DB.Where("email = ?", email).First(&user).Error
}

func (u *User) CreateUser(user User) error {
	return dao.DB.Create(&user).Error
}
