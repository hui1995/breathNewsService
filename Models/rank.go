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
 * @Date: 2020/2/26 3:56 PM
 */

type Rank struct {
	gorm.Model

	UserId   int
	UserName string
	IsRobot  int
	Score    int
	Position int
}

//func init() {
//	table := Mysql.DB.HasTable(Rank{})
//	if !table {
//		Mysql.DB.CreateTable(Rank{})
//	}
//}
func (this *Rank) GetRank() []Rank {
	var ranks []Rank
	Mysql.DB.Order("position").Limit(100).Find(&ranks)
	return ranks
}

func (this *Rank) GetCurrentRank(userId int) Rank {
	var rank Rank
	Mysql.DB.Where("user_id = ?", userId).First(&rank)
	return rank
}
