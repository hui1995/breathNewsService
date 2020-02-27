package Models

import (
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
}

//func init() {
//	table := Mysql.DB.HasTable(Rank{})
//	if !table {
//		Mysql.DB.CreateTable(Rank{})
//	}
//}
