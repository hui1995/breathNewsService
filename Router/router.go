package Router

import (
	"breathNewsService/Controllers"
	"breathNewsService/Middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	//router.LoadHTMLGlob("templates/**")

	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.Cors())
	router.Use(Middlewares.GetUserId())

	// 使用 session(cookie-based)
	router.POST("/login/check", Controllers.Login)
	router.GET("home/channel", Controllers.ArticleList)
	router.POST("load/config", Controllers.LoadsConfig)

	v1 := router.Group("/api")
	v1.Use(Middlewares.LoginAuth())
	{
		v1.GET("/article/:id", Controllers.ArticleDetail)
		v1.GET("/auth/article/", Controllers.AuthPermissionArticle)

	}

	router.Run(":8080")
}
