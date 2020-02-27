package Models

import (
	"breathNewsService/Databases/Mysql"
	"github.com/jinzhu/gorm"
)

type CommonConfig struct {
	gorm.Model

	Group  string
	Key    string
	Value  string
	remark string
}

func init() {
	table := Mysql.DB.HasTable(CommonConfig{})
	if !table {
		Mysql.DB.CreateTable(CommonConfig{})
	}
}
