package Controllers

import (
	"breathNewsService/Controllers/Request"
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
func LoadsConfig(c *gin.Context) {
	var request Request.DevideInfoRequest

	if c.BindJSON(&request) == nil {

		config := Services.CheckPlatform(request.DevideId, request.Platform, request.Manufacturer, request.Model)
		c.JSON(http.StatusOK,
			gin.H{"code": 1, "data": config})
	} else {
		c.JSON(http.StatusOK,
			gin.H{"code": 1, "message": "参数错误"})
	}

}
