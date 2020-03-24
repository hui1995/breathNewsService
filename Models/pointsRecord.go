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
 * @Date: 2020/3/24 9:22 PM
 */
type PointsRecord struct {
	gorm.Model

	UserId int
	Day    string
	Count  int
}

func init() {
	table := Mysql.DB.HasTable(PointsRecord{})
	if !table {
		Mysql.DB.CreateTable(PointsRecord{})
	}
}
func (this *PointsRecord) InsertPointsRAecord(userId int, day string) PointsRecord {
	var pointsPaecord PointsRecord
	Mysql.DB.Where(PointsRecord{UserId: userId, Day: day}).FirstOrCreate(&pointsPaecord)
	return pointsPaecord

}
func (this *PointsRecord) UpdatePointsRecord(userId int, day string, count int) bool {

	var pointsRecord PointsRecord

	Mysql.DB.Where(&PointsRecord{UserId: userId}).First(&pointsRecord)

	Mysql.DB.Model(&pointsRecord).Update("Count", count)
	return true
}
