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

	var userHandler handler.UserHandler = handler.UserHandler{}
	var postHandler handler.PostHandler = handler.PostHandler{}
	var commentHandler handler.CommentHandler = handler.CommentHandler{}

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", userHandler.RegistHandler)
		userGroup.POST("/login", userHandler.LoginHandler)
	}

	postGroup := r.Group("/post")
	{
		postGroup.POST("/insertPost", middleware.AuthMiddleware(), postHandler.InsertPostHandler)
		postGroup.GET("/getPostById/:id", postHandler.GetPostByIdHandler)
		postGroup.GET("/getAllPost", postHandler.GetAllPostHandler)
		postGroup.PUT("/updatePost", middleware.AuthMiddleware(), postHandler.UpdatePostHandler)
		postGroup.DELETE("/deletePostById/:id", middleware.AuthMiddleware(), postHandler.DeletePostByIdHandler)
	}

	commentGroup := r.Group("/comment")
	{
		commentGroup.POST("/insertCommnet", middleware.AuthMiddleware(), commentHandler.CreateCommnetHandler)
		commentGroup.GET("/getCommnetByPostId/:id", commentHandler.GetCommentByPostIdHandler)
	}

	r.Run(":" + config.Server.Port)
}
