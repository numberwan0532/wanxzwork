package api

import (
	"github.com/gin-gonic/gin"
	"github.com/numberwan0532/wanxzwork/task4/configs"
	"github.com/numberwan0532/wanxzwork/task4/internal/handler"
	"github.com/numberwan0532/wanxzwork/task4/internal/middleware"
	"github.com/sirupsen/logrus"
)

func Start(config *configs.Config, appLog *logrus.Logger) {
	r := gin.Default()
	r.Use(middleware.LoggingMiddleware(appLog))
	r.Use(middleware.ErrorMiddleware(appLog))

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handler.RegistHandler)
		userGroup.POST("/login", handler.LoginHandler)
	}

	postGroup := r.Group("/post")
	{
		postGroup.POST("/insertPost", middleware.AuthMiddleware(), handler.InsertPostHandler)
		postGroup.GET("/getPostById/:id", handler.GetPostByIdHandler)
		postGroup.GET("/getAllPost", handler.GetAllPostHandler)
		postGroup.PUT("/updatePost", middleware.AuthMiddleware(), handler.UpdatePostHandler)
		postGroup.DELETE("/deletePostById/:id", middleware.AuthMiddleware(), handler.DeletePostByIdHandler)
	}

	commentGroup := r.Group("/comment")
	{
		commentGroup.POST("/insertCommnet", middleware.AuthMiddleware(), handler.CreateCommnetHandler)
		commentGroup.GET("/getCommnetByPostId/:id", handler.GetCommentByPostIdHandler)
	}

	r.Run(":" + config.Server.Port)
}
