package Controllers

import (
	"breathNewsService/Services"
	"github.com/gin-gonic/gin"
	"html/template"
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

		c.HTML(http.StatusOK, "detail.html", gin.H{
			"code":        1,
			"message":     "获取成功",
			"data":        &article,
			"content":     template.HTML(article.Content),
			"create_time": article.CreatedAt.Format("2006年01月02日"),
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
	info := Services.GetMain(userID)

	articlelst := Services.GetHomeArticleList(channelInt, userID)
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "获取成功",
		"data":    articlelst,
		"main":    info,
	})
	return
}
