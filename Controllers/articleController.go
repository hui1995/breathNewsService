package Controllers

import (
	"breathNewsService/Models"
	"fmt"
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
 * @Date: 2020/2/27 2:55 PM
 */

func ArticleDetail(c *gin.Context) {

	var articleModel Models.Article
	id:=c.Param("id")

	idInt,_:=strconv.Atoi(id)

	article, err := articleModel.FindById(idInt)
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": err,
		"data":    &article,
	})



}
func ArticleList(c *gin.Context){
	page,_ :=strconv.Atoi(c.Query("page"))
	pageSize,_:=strconv.Atoi(c.Query("pageSize"))
	fmt.Println(page)
	fmt.Println(pageSize)
	var channelInt int;
	channelInt=0
	if channel,isExist :=c.GetQuery("channel");isExist==true{
		channelInt,_=strconv.Atoi(channel)

	}
	var articleModel Models.Article
	articlelst:=articleModel.FindByChannel(channelInt,page,pageSize)
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": "获取成功",
		"data":    articlelst,
	})
	return
}
