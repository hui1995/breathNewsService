package Models

import (
	"github.com/jinzhu/gorm"
)

type Channel struct {
	gorm.Model

	Name     string
	Desction string
}

//func init() {
//	table := Mysql.DB.HasTable(Channel{})
//	if !table {
//		Mysql.DB.CreateTable(Channel{})
//	}
//}
