package Models

import (
	"github.com/jinzhu/gorm"
)

//收入记录表

type IncomeRecord struct {
	gorm.Model
	Type   int
	Money  float64
	userId int
}

//func init() {
//	table := Mysql.DB.HasTable(IncomeRecord{})
//	if !table {
//		Mysql.DB.CreateTable(IncomeRecord{})
//	}
//}
