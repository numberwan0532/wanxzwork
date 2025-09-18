package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/numberwan0532/wanxzwork/task4/pkg/globalerrors"
)

var secretKey string = os.Getenv("secretKey")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if len(tokenString) < 7 {
			c.AbortWithStatusJSON(http.StatusOK, globalerrors.Response{
				Code:    401,
				Message: "未提供认证Token",
			})
			return
		}

		token, err := jwt.Parse(tokenString[7:], func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, globalerrors.Response{
				Code:    401,
				Message: "Token校验失败",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 验证标准声明
			if exp, ok := claims["exp"].(float64); !ok || time.Now().Unix() > int64(exp) {
				c.AbortWithStatusJSON(http.StatusOK, globalerrors.Response{
					Code:    401,
					Message: "Token已过期",
				})
				return
			}
			// 提取自定义字段
			userIDFloat := claims["userID"].(float64)
			c.Set("userID", uint(userIDFloat))
			c.Set("username", claims["username"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusOK, globalerrors.Response{
				Code:    401,
				Message: "无效Token",
			})
		}
	}
}
