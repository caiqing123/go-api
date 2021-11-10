package routes

import (
	"api/controllers"
	_ "api/docs"
	"api/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRoutes(router *gin.Engine) {
	router.Use(gin.Recovery()) // error handle

	v1 := router.Group("/v1")
	{
		v1.GET("api",
			middleware.CorsMiddleware(),
			func(ctx *gin.Context) {
				hello := controllers.ApiController{}
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

	pay := router.Group("/pay", middleware.CorsMiddleware())
	{
		wechat := pay.Group("wechat")
		{
			wechat.POST("/v2",
				func(ctx *gin.Context) {
					wechat := controllers.WechatV2PayController{}
					wechat.Index(ctx)
				},
			)

			wechat.Any("/v2/notify",
				func(ctx *gin.Context) {
					wechat := controllers.WechatV2PayController{}
					wechat.Notify(ctx)
				},
			)

			wechat.POST("/v3",
				func(ctx *gin.Context) {
					wechat := controllers.WechatV3PayController{}
					wechat.Index(ctx)
				},
			)
		}

		alipay := pay.Group("alipay")
		{
			alipay.POST("/",
				func(ctx *gin.Context) {
					alipay := controllers.AlipayController{}
					alipay.Index(ctx)
				},
			)
			alipay.Any("/notify",
				func(ctx *gin.Context) {
					alipay := controllers.AlipayController{}
					alipay.Notify(ctx)
				},
			)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
