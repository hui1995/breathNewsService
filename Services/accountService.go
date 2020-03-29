package Services

import (
	"breathNewsService/Controllers/Response"
	"breathNewsService/Models"
	jwtgo "github.com/dgrijalva/jwt-go"
	"strconv"

	myjwt "breathNewsService/Middlewares"
	"breathNewsService/utils"
	"log"

	"time"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/5 3:51 PM
 */

func generateToken(user Models.User) (string, error) {
	j := &myjwt.JWT{
		[]byte("NamelyThinking"),
	}
	claims := myjwt.CustomClaims{
		user.RealNum,
		user.UserName,
		user.Phone,
		true,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),   // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 360000), // 过期时间 一小时
			Issuer:    "breathCoder",                     //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {

		return "", err
	}

	log.Println(token)

	return token, nil
}

func CheckLogin(phone, password string) (int, *Response.AccountResponse) {

	var userModel Models.User
	user, err := userModel.FindByPhone(phone)
	if err != nil {
		return 0, nil

	} else {

		if user.Status != 0 && user.Status != 1 {
			return -1, nil
		}

		if utils.MD5(password) != user.Password {
			return -2, nil
		}
		var accountResponse Response.AccountResponse
		accountResponse.UserName = user.UserName
		accountResponse.Money = user.Money
		accountResponse.ActiveMoney = user.ActiveMoney
		accountResponse.FreezeMoney = user.FreezeMoney
		accountResponse.RealNum = user.RealNum

		token, err := generateToken(*user)
		if err != nil {
			return -4, nil
		}
		accountResponse.Token = token

		return 1, &accountResponse

	}

}

func UpdateAliPay(userId int, alipay, userName string) bool {

	var user Models.User
	if !user.SelectByAliPay(alipay) {
		return false

	}
	return user.UpdateAlipay(userId, alipay, userName)
}

type UserInfo struct {
	UserName string
	Money    float64
	Point    int
}

func GetUserInfo(userId int) UserInfo {
	var user Models.User
	var userPoint Models.UserPoints
	userCurrent, _ := user.FindByRealNum(userId)
	userPoint = userPoint.SelectPointsByUserId(userId)
	var userInfo UserInfo
	userInfo.Money = userCurrent.Money
	userInfo.UserName = userCurrent.UserName
	userInfo.Point = userPoint.Points
	return userInfo

}

func Invite(inviter, invitee int) bool {
	var invitte Models.Invite
	if invitte.FindeByInvitee(invitee) {
		invitte.InsertInvite(inviter, invitee)
		return true

	}
	return false
}
func InviteeList(userId int) []Models.Invite {
	var invitte Models.Invite
	return invitte.FIndInviteeByInviter(userId)
}

type AnyInvite struct {
	Count   int
	Url     string
	Inviter int
}

func InviteAny(userId int) AnyInvite {

	var invitee Models.Invite
	count := invitee.FindCountByInviter(userId)
	var anyInvite AnyInvite
	anyInvite.Url = "https://atvspn.jmlk.co/AA4o?mw_dynp_u_id=" + strconv.Itoa(userId)
	anyInvite.Count = count
	anyInvite.Inviter = userId

	return anyInvite
}
