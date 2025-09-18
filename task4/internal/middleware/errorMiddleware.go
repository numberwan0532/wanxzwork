package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/numberwan0532/wanxzwork/task4/pkg/globalerrors"
	"github.com/sirupsen/logrus"
)

// 统一错误处理中间件
func ErrorMiddleware(appLog *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				appLog.WithFields(logrus.Fields{
					"path":  c.Request.URL.Path,
					"error": err.Error(),
					"stack": fmt.Sprintf("%+v", err),
				}).Error("请求处理异常")
			}
			err := c.Errors.Last()
			c.JSON(http.StatusOK, globalerrors.Response{
				Code:    500,
				Message: err.Error(),
			})
		}
	}
}
