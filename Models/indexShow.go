package Models

import (
	"github.com/jinzhu/gorm"
)

type IndexShow struct {
	gorm.Model
	ChannelId int
}

//func init() {
//	table := Mysql.DB.HasTable(IndexShow{})
//	if !table {
//		Mysql.DB.CreateTable(IndexShow{})
//	}
//}
