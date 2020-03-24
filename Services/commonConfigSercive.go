package Services

import (
	"breathNewsService/Models"
	"encoding/json"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"strconv"

	myjwt "breathNewsService/Middlewares"
	"time"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/10 10:08 PM
 */

type PlatformConfig struct {
	Platform string
	Open     bool
}

type AdsConfig struct {
	Platform string
	Ads      []string
}

type ConfigInfo struct {
	IsOpen   bool
	Rate     int
	Ads      []string
	Limit    int
	Interval int
	UserId   int
	Points   int
	Message  string
	IsFirst  bool
	Token    string
}

func generateTokenByDevide(realNum int, username string, phone string) (string, error) {
	j := &myjwt.JWT{
		[]byte("NamelyThinking"),
	}
	claims := myjwt.CustomClaims{
		realNum,
		username,
		phone,
		true,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),    // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600000), // 过期时间 一小时
			Issuer:    "breathCoder",                      //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {

		return "", err
	}

	return token, nil
}

func FirstRegister(devideId string, platform string, manufacturer string, model string, points int) (bool, int, string) {
	var devideInfo Models.DevideInfo
	var user Models.User
	var userPoints Models.UserPoints
	if !devideInfo.IsExistByDevideId(devideId) {
		realId := user.FindLastUser()
		realId = realId + 1
		user.InsertInfo(realId, 0, "用户"+strconv.Itoa(realId))
		devideInfo.InsertInfo(realId, devideId, platform, manufacturer, model)
		userPoints.InsertInfo(realId, points)

		token, err := generateTokenByDevide(realId, "用户"+strconv.Itoa(realId), "")
		if err == nil {
			return true, realId, token
		}
		return true, realId, ""

	} else {
		info, err := devideInfo.FindByDevideId(devideId)
		if err == nil {
			realId := info.UserId
			token, err := generateTokenByDevide(realId, "用户"+strconv.Itoa(realId), "")
			if err == nil {
				return false, realId, token
			}
			return false, realId, ""

		}
	}
	return false, 0, ""
}

func CheckPlatform(devideId string, platform string, manufacturer string, model string) ConfigInfo {

	var commconfig Models.CommonConfig
	var configInfo ConfigInfo

	//第一次注册获取注册信息
	points := 66

	isFisrt, userId, token := FirstRegister(devideId, platform, manufacturer, model, points)

	value := commconfig.FindByGroupAndKey("platform", "version")
	var platforms []PlatformConfig
	err := json.Unmarshal([]byte(value), &platforms)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)

	}

	configInfo.IsOpen = false
	configInfo.Token = token

	for _, item := range platforms {

		if item.Platform == platform {
			configInfo.IsOpen = true
			break

		}

	}
	if isFisrt {
		configInfo.UserId = userId
		configInfo.Points = points
		configInfo.IsFirst = isFisrt
		if configInfo.IsOpen {
			configInfo.Message = "欢迎您第" + strconv.Itoa(userId) + "位用户，首次注册，赠送您" + strconv.Itoa(points) + "J币"
		} else {
			configInfo.Message = "欢迎您第使用"

		}
	}

	//loads ads platform config
	value = commconfig.FindByGroupAndKey("platform", "ad")

	var adsConfig []AdsConfig

	err = json.Unmarshal([]byte(value), &adsConfig)

	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)

	} else {
		for _, item := range adsConfig {

			if item.Platform == platform {
				configInfo.Ads = item.Ads

			}
		}
	}
	value = commconfig.FindByGroupAndKey("ad", "rate")
	rate, err := strconv.Atoi(value)
	if err != nil {
		rate = 20
	}
	configInfo.Rate = rate

	value = commconfig.FindByGroupAndKey("ad", "queryRule")
	var queryRule []QueryRule
	err3 := json.Unmarshal([]byte(value), &queryRule)
	if err3 != nil {
		configInfo.Limit = 20
		configInfo.Interval = 20

	} else {
		configInfo.Limit = queryRule[0].Limit
		configInfo.Interval = queryRule[0].Interval
	}

	fmt.Println(configInfo)

	return configInfo
}
