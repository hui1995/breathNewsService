package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/21 4:54 PM
 */
func PrivateController(c *gin.Context) {

	c.HTML(http.StatusOK, "private.html", nil)
}

func Contractontroller(c *gin.Context) {

	c.HTML(http.StatusOK, "contract.html", nil)
}
func RankHelpController(c *gin.Context) {

	c.HTML(http.StatusOK, "rank_help.html", nil)
}

func InviteController(c *gin.Context) {
	c.HTML(http.StatusOK, "invite.html", nil)
}

func Test(c *gin.Context) {
	c.HTML(http.StatusOK, "invitelst.html", nil)
}
