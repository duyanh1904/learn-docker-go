package server

import (
	"github.com/duyanh1904/learn-docker-go/controllers"
	"github.com/duyanh1904/learn-docker-go/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	post := new(controllers.PostController)

	router.Use(middlewares.AuthMiddleware())
	router.GET("/health", health.Status)
	router.POST("/add", post.CreatePost)
	router.GET("/kafka", post.RunKafka)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retrieve)
		}
	}
	return router

}
