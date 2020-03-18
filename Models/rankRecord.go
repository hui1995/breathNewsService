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
 * @Date: 2020/3/16 11:47 PM
 */
type RankRecord struct {
	gorm.Model

	UserId   int
	UserName string
	IsRobot  int
	Score    int
	Reward   float64
	Position int
	Day      int
}

//func init() {
//	table := Mysql.DB.HasTable(RankRecord{})
//	if !table {
//		Mysql.DB.CreateTable(RankRecord{})
//	}
//}

func (this *RankRecord) GetYesRank(day string) []RankRecord {

	var ranks []RankRecord
	Mysql.DB.Where("day = ?", day).Order("position").Limit(10).Find(&ranks)
	return ranks

}

func (this *RankRecord) GetCurrentRank(day string, userId int) RankRecord {
	var rank RankRecord
	Mysql.DB.Where("day = ? and user_id = ?", day, userId).First(&rank)
	return rank
}
