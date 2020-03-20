package Controllers

import (
	"breathNewsService/Services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/17 11:36 PM
 */
func GetTodayRank(c *gin.Context) {

	userID := c.GetInt("userId")
	if userID <= 100000 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "该用户不存在",
		})
		return
	}

}

func GetYestodayRank(c *gin.Context) {

	userID := c.GetInt("userId")
	if userID <= 100000 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "该用户不存在",
		})
		return
	}
	data := Services.GetYesRank(userID)
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"data":    data,
		"message": "success",
	})

}
func TestH5(c *gin.Context) {
	data := Services.GetTodayRank(888888)
	fmt.Println(data)
	c.HTML(http.StatusOK, "rank.html", gin.H{
		"title": "今日榜单", "data": data,
	})
}
