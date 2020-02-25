package Controllers

import (
	"breathNewsService/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestInsert(c *gin.Context) {
	var testService Services.Test

	err := c.ShouldBindJSON(&testService)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"message": "Insert() error!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "success",
		"data": "s",
	})

}