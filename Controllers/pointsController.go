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
	userID := c.GetInt("userId")

	points := Services.AddPoints(userID)
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "获取成功",
		"data":    points,
	})
}
func AuthPointDraw(c *gin.Context) {
	userID := c.GetInt("userId")
	count, invert := Services.SelectLimit(userID)
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "今日领取已经达到上限",
		})
	} else if count == -1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "领取太频繁，请" + strconv.Itoa(invert) + "稍后再试",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "可以抽取",
			"data":    count,
		})
	}

}

func DrawReadReal(c *gin.Context) {
	userID := c.GetInt("userId")
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	Services.ReadIsReal(userID, idInt)
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "成功",
	})

}
