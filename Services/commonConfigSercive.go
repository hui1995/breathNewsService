package Services

import (
	"breathNewsService/Models"
	"encoding/json"
	"fmt"
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

func CheckPlatform(platform string) bool {

	var commconfig Models.CommonConfig

	value := commconfig.FindByGroupAndKey("platform", "version")
	var platforms []PlatformConfig
	err := json.Unmarshal([]byte(value), &platforms)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return false
	}

	for _, item := range platforms {

		if item.Platform == platform {
			return item.Open

		}

	}
	return false
}
