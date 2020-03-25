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
 * @Description: * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/2/27 2:55 PM
 */

/**
获取详情
*/

func ArticleDetail(c *gin.Context) {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	article, err := Services.GetArticleDeatil(uint(idInt))
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "获取成功",
			"data":    &article,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err,
		})

	}

}

/**
获取列表
*/
func ArticleList(c *gin.Context) {

	userID := c.GetInt("userId")

	var channelInt int
	channelInt = 0
	if channel, isExist := c.GetQuery("channel"); isExist == true {
		channelInt, _ = strconv.Atoi(channel)

	}
	articlelst := Services.GetHomeArticleList(channelInt, userID)
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "获取成功",
		"data":    articlelst,
	})
	return
}

func GetMainInfo(c *gin.Context) {
	userID := c.GetInt("userId")

	info := Services.GetMain(userID)
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "获取成功",
		"data":    info,
	})
}
