package Controllers

import (
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
 * @Date: 2020/3/21 5:08 PM
 */
func GetInComeRecord(c *gin.Context) {
	userID := c.GetInt("userId")

	reuslt := Services.GetIncomeRecord(userID)
	c.HTML(http.StatusOK, "history.html", gin.H{"data": reuslt})

}

func IncomeByDetail(c *gin.Context) {
	userID := c.GetInt("userId")
	Services.AddIncome(userID)
	c.JSON(http.StatusOK,
		gin.H{
			"code": 1, "msg": "获取成功"})

}
