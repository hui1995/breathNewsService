package Controllers

import (
	"breathNewsService/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/12 12:58 AM


 */
func GetPlatformVersion(c *gin.Context) {
	platform := c.Query("platform")

	isOpen := Services.CheckPlatform(platform)

	c.JSON(http.StatusOK,
		gin.H{"code": 1, "isOpen": isOpen})

}
