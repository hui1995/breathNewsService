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
	// 使用 session(cookie-based)
	//router.Use(sessions.Sessions("myyyyysession", Sessions.Store))
	//v1 := router.Group("v1")
	//{
	router.GET("/article", Controllers.ArticleDetail)
	//}

	router.Run(":8080")
}
