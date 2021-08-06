package routes

import (
	"api/controllers"
	"api/middleware"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	router.Use(gin.Recovery()) // error handle

	v1 := router.Group("/v1")
	{
		v1.GET("hello",
			middleware.CorsMiddleware(),
			func(ctx *gin.Context) {
				hello := controllers.HelloController{}
				hello.Index(ctx)
			},
		)

		v1.POST("users/add",
			middleware.AuthMiddleware(),
			func(ctx *gin.Context) {
				hello := controllers.UserController{}
				hello.Add(ctx)
			},
		)

		v1.POST("auth", func(ctx *gin.Context) {
			auth := controllers.AuthController{}
			auth.Index(ctx)
		})

		v1.POST("test", func(ctx *gin.Context) {
			test := controllers.TestController{}
			test.Index(ctx)
		})
	}
}
