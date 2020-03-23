package Controllers

import (
	"breathNewsService/Services"
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
 * @Date: 2020/3/21 9:13 PM
 */
func ProductController(c *gin.Context) {

	userID := c.GetInt("userId")
	data := Services.FindProductInfo(userID)
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"data":    data,
		"message": "获取成功",
	})
}

func OrderProduct(c *gin.Context) {
	userID := c.GetInt("userId")

	price := c.Query("price")
	priceInt, _ := strconv.ParseFloat(price, 64)

	if Services.AddOrderInfo(userID, priceInt) {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "下订单成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "订单过多",
		})
	}

}
func OrderList(c *gin.Context) {
	userID := c.GetInt("userId")
	data := Services.GetOrderList(userID)
	fmt.Println(data)
	c.HTML(http.StatusOK, "orderhistory.html", gin.H{"data": data})

}
