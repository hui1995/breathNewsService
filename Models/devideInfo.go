package Models

import (
	"breathNewsService/Databases/Mysql"
	"github.com/jinzhu/gorm"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/14 9:06 PM
 */
type DevideInfo struct {
	gorm.Model

	UserId   int
	DevideId string
	// 密码
	Platform     string
	Manufacturer string
	ModelM       string
}

func (this *DevideInfo) IsExistByDevideId(devideId string) bool {
	var count int
	Mysql.DB.Model(&DevideInfo{}).Where("devide_id = ?", devideId).Count(&count)
	if count == 0 {

		return false

	}

	return true
}

func (this *DevideInfo) InsertInfo(userId int, devideId string, platform string, manufacturer string, model string) bool {
	var incomeRecord = DevideInfo{UserId: userId, DevideId: devideId, Platform: platform, Manufacturer: manufacturer, ModelM: model}
	Mysql.DB.Create(&incomeRecord)
	return true
}

func init() {
	table := Mysql.DB.HasTable(DevideInfo{})
	if !table {
		Mysql.DB.CreateTable(DevideInfo{})
	}
}
