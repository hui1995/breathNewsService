package Controllers

import (
	"breathNewsService/Controllers/Request"
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
 * @Date: 2020/2/27 9:33 PM
 */
func Login(c *gin.Context) {
	var request Request.LoginRequest
	if c.BindJSON(&request) == nil {

		//	说明数据正确
		code, userInfo := Services.CheckLogin(request.Phone, request.Pwd)
		if code == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "该手机号未注册"})
			return
		} else if code == -1 {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "该用户非正常用户"})
			return
		} else if code == -2 {
			c.JSON(http.StatusOK, gin.H{
				"status": -2,
				"msg":    "密码不正确"})
			return
		} else if code == -4 {
			c.JSON(http.StatusOK, gin.H{
				"status": -4,
				"msg":    "生成token失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "登录成功",
			"data":   userInfo})
		return

	} else {
		fmt.Println(c.BindJSON(&request))

		c.JSON(http.StatusOK, gin.H{
			"status": -3,
			"msg":    "参数错误"})
		return
	}

}

func UpdateAlipay(c *gin.Context) {
	userID := c.GetInt("userId")

	var request Request.AlipayRequest
	if c.BindJSON(&request) == nil {

		if Services.UpdateAliPay(userID, request.Alipay, request.RealName) {
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    "添加成功"})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "该账号绑定账号过多"})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -3,
			"msg":    "参数错误"})
		return
	}

}

func UserInfoDetail(c *gin.Context) {

	userID := c.GetInt("userId")
	userInfo := Services.GetUserInfo(userID)
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"data":   userInfo,
		"msg":    "获取成功"})

}
func InviteInfoController(c *gin.Context) {
	userId := c.GetInt("userId")
	inviter := c.Param("inviter")
	idInt, _ := strconv.Atoi(inviter)
	Services.Invite(idInt, userId)
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "邀请成功",
	})
}

func InviteH5Controller(c *gin.Context) {
	userID := c.GetInt("userId")

	data := Services.InviteAny(userID)

	c.HTML(http.StatusOK, "invite.html", gin.H{
		"data": data,
	})
}

func InviteHistoryController(c *gin.Context) {
	userID := c.Param("userId")
	idInt, _ := strconv.Atoi(userID)
	lst := Services.InviteeList(idInt)

	c.HTML(http.StatusOK, "invitelst.html", gin.H{
		"data": lst,
	})

}
