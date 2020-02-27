package Controllers

import (
	"breathNewsService/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/2/27 2:55 PM
 */

func ArticleDetail(c *gin.Context) {

	var articleModel Models.Article

	article, err := articleModel.FindById(1)

	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": err,
		"data":    article,
	})
	return

}
