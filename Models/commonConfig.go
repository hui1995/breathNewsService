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
func (this *CommonConfig) FindByGroupAndKey(group, key string) string {

	var commonConfig CommonConfig
	Mysql.DB.Where(&CommonConfig{Group: group, Key: key}).First(&commonConfig)
	value := commonConfig.Value
	return value

}
