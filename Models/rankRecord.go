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
	Days     int
}

func init() {
	table := Mysql.DB.HasTable(RankRecord{})
	if !table {
		Mysql.DB.CreateTable(RankRecord{})
	}
}
