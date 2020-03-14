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
 * @Date: 2020/2/27 2:28 PM
 */

type UserPoints struct {
	gorm.Model

	UserId int
	Points int
}

//func init() {
//	table := Mysql.DB.HasTable(UserPoints{})
//	if !table {
//		Mysql.DB.CreateTable(UserPoints{})
//	}
//}

// 搜索用户积分
func (this *UserPoints) SelectPointsByUserId(userId int) UserPoints {
	var userPoints UserPoints

	Mysql.DB.Where(&UserPoints{UserId: userId}).First(&userPoints)
	return userPoints

}

//减去积分
func (this *UserPoints) SubPoints(userId, points int) bool {
	var userPoints UserPoints

	Mysql.DB.Where(&UserPoints{UserId: userId}).First(&userPoints)

	var pointsNow = userPoints.Points - points
	Mysql.DB.Model(&userPoints).Update("points", pointsNow)
	return true

}
func (this *UserPoints) InsertInfo(realNum, points int) bool {
	user := UserPoints{UserId: realNum, Points: points}

	Mysql.DB.Create(&user)
	return true

}
