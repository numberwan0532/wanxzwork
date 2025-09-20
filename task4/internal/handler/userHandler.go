package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/numberwan0532/wanxzwork/task4/internal/service"
	"github.com/numberwan0532/wanxzwork/task4/pkg/globalerrors"
)

type UserHandler struct {
}

var userServce service.UserService = service.UserService{}

func (userHandler *UserHandler) RegistHandler(c *gin.Context) {
	err := userServce.Register(c)
	CommonReturn(c, nil, err)
}

func (userHandler *UserHandler) LoginHandler(c *gin.Context) {
	data, err := userServce.Login(c)
	CommonReturn(c, data, err)
}

func CommonReturn(c *gin.Context, data interface{}, err error) {
	if err != nil {
		// c.JSON(http.StatusOK, gin.H{
		// 	"code": "00001", "msg": err.Error()})
		c.Error(err)
	} else {
		c.JSON(http.StatusOK,
			globalerrors.Response{
				Code:    0,
				Message: "success",
				Data:    data,
			})
	}
}
