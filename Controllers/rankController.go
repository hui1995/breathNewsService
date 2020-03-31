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
 * @Date: 2020/3/17 11:36 PM
 */
func GetTodayRank(c *gin.Context) {
	var rankTypeInt int
	userID := c.GetInt("userId")

	if rankType, isExist := c.GetQuery("type"); isExist == true {
		rankTypeInt, _ = strconv.Atoi(rankType)

	} else {
		rankTypeInt = 0
	}
	if rankTypeInt == 0 {
		data, meData := Services.GetTodayRank(userID)
		c.HTML(http.StatusOK, "rank.html", gin.H{
			"title": "今日榜单", "data": data, "meData": meData, "value": "阅读指数",
		})

	} else {

		data, meData := Services.GetYesRank(userID)
		c.HTML(http.StatusOK, "rank3.html", gin.H{
			"title": "昨日榜单", "data": data, "meData": meData, "value": "奖励",
		})

	}

}
