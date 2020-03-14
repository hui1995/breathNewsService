package Services

import (
	"breathNewsService/Models"
	"encoding/json"
	"fmt"
	"strconv"
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

type QueryRule struct {
	Limit    int
	Interval int
}

type ConfigInfo struct {
	IsOpen   bool
	Rate     int
	Ads      []string
	Limit    int
	Interval int
	userId   int
}

func FirstRegister(devideId string, platform string, manufacturer string, model string) int {
	var devideInfo Models.DevideInfo
	var user Models.User
	var userPoints Models.UserPoints
	if !devideInfo.IsExistByDevideId(devideId) {
		realId := user.FindLastUser()
		realId = realId + 1
		user.InsertInfo(realId, 0, "用户"+strconv.Itoa(realId))
		devideInfo.InsertInfo(realId, devideId, platform, manufacturer, model)
		userPoints.InsertInfo(realId, 66)
		return realId - 100000

	}
	return 0
}

func CheckPlatform(devideId string, platform string, manufacturer string, model string) ConfigInfo {

	var commconfig Models.CommonConfig
	var configInfo ConfigInfo

	//第一次注册获取注册信息

	userId := FirstRegister(devideId, platform, manufacturer, model)
	if userId != 0 {
		configInfo.userId = userId
	}
	value := commconfig.FindByGroupAndKey("platform", "version")
	var platforms []PlatformConfig
	err := json.Unmarshal([]byte(value), &platforms)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)

	}

	configInfo.IsOpen = false

	for _, item := range platforms {

		if item.Platform == platform {
			configInfo.IsOpen = true
			break

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
				fmt.Println(item.Ads)
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