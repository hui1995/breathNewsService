package Router

import (
	"breathNewsService/Controllers"
	"breathNewsService/Middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**")

	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.Cors())
	router.Use(Middlewares.GetUserId())

	// 使用 session(cookie-based)
	router.POST("/login/check", Controllers.Login)
	router.GET("home/channel", Controllers.ArticleList)
	router.POST("load/config", Controllers.LoadsConfig)
	router.GET("/rank/", Controllers.GetTodayRank)
	router.GET("/private/", Controllers.PrivateController)
	router.GET("/contract/", Controllers.Contractontroller)
	router.GET("/rankhelp", Controllers.RankHelpController)
	router.GET("/main/info", Controllers.GetMainInfo)

	v1 := router.Group("/api")
	v1.Use(Middlewares.LoginAuth())
	{
		v1.GET("/article/:id", Controllers.ArticleDetail)
		v1.GET("/history/", Controllers.GetInComeRecord)
		v1.GET("/auth/article/", Controllers.AuthPermissionArticle)
		v1.GET("/product", Controllers.ProductController)
		v1.GET("/order", Controllers.OrderProduct)
		v1.POST("/alipay", Controllers.UpdateAlipay)
		v1.GET("/userinfo", Controllers.UserInfoDetail)
		v1.GET("/orderlst", Controllers.OrderList)
		v1.GET("/invite", Controllers.InviteController)
		v1.GET("/add/points", Controllers.AddPoints)
		v1.GET("/auth/draw", Controllers.AuthPointDraw)

	}

	router.Run(":80")
}
