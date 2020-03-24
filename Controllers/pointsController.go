package Controllers

import (
	"breathNewsService/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/9 12:55 AM
 */

func AuthPermissionArticle(c *gin.Context) {

	userID := c.GetInt("userId")
	articleId := c.Query("id")
	idInt, _ := strconv.Atoi(articleId)

	points, err := Services.Consumption(uint(idInt), userID)
	if err == -1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "获取文章错误",
		})

	} else if err == -2 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -2,
			"message": "积分不足",
		})

	} else if err == -3 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -3,
			"message": "未获得奖励",
		})
	} else if err == -4 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -4,
			"message": "减去积分错误",
		})

	} else if err == -5 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -5,
			"message": "已经查看过",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "获取成功",
			"data":    points,
		})
	}

}
func AddPoints(c *gin.Context) {
	userID := c.Query("user_id")
	user_id, _ := strconv.Atoi(userID)
	isTrue := Services.AddPoints(user_id)
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "获取成功",
		"data":    isTrue,
	})
}
