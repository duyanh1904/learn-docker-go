package server

import (
	"github.com/duyanh1904/learn-docker-go/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/santosh/gingo/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	post := new(controllers.PostController)
	user := new(controllers.UserController)
	grpc := new(controllers.GrpcController)
	curl := new(controllers.CurlController)

	//router.Use(middlewares.AuthMiddleware())
	router.GET("/health", health.Status)
	router.GET("/JSON", health.JsonArrays)
	router.GET("/channel", health.MakeChannel)
	router.POST("/add", post.CreatePost)
	router.GET("/kafka", post.RunKafka)
	router.GET("/users", user.Retrieve)
	router.POST("/update", user.Update)
	router.GET("/get-token", user.GenToken)
	router.GET("/validate", post.TestValidateApi)
	router.GET("/grpc", grpc.GetPerson)
	router.GET("/get_list_partner", curl.GetListPartner)

	//swag
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
