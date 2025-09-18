package service

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/numberwan0532/wanxzwork/task4/internal/model"
	"golang.org/x/crypto/bcrypt"
)

var secretKey string = os.Getenv("secretKey")

func Register(c *gin.Context) error {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	_, err1 := model.GetUserByUsername(user.Username)
	if err1 == nil {
		return errors.New("该用户名已存在")
	}
	err2 := model.GetUserByEmail(user.Email)
	if err2 == nil {
		return errors.New("该邮箱已存在")
	}
	return model.CreateUser(user)
}

func Login(c *gin.Context) (map[string]string, error) {
	var marr = make(map[string]string)
	var user model.User
	c.ShouldBindJSON(&user)
	userOld, err := model.GetUserByUsername(user.Username)
	if err != nil {
		return marr, errors.New("该用户不存在")
	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(userOld.Password), []byte(user.Password)); err != nil {
		return marr, errors.New("密码错误")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   userOld.ID,
		"username": userOld.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	if len(secretKey) == 0 {
		return marr, errors.New("token 密钥为空")
	}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return marr, errors.New("failed to generate token")
	}
	marr["username"] = userOld.Username
	marr["token"] = tokenString
	return marr, nil
}
