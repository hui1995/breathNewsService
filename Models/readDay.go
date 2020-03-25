package Models

import (
	"breathNewsService/Databases/Mysql"
	"fmt"
	"github.com/jinzhu/gorm"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/25 5:21 PM
 */

//连续阅读天数
type ReadRecord struct {
	gorm.Model

	UserId int
	Day    string
	Count  int
}

func init() {
	table := Mysql.DB.HasTable(ReadRecord{})
	if !table {
		Mysql.DB.CreateTable(ReadRecord{})
	}
}

//获取或者创建
func (this *ReadRecord) GetRankRecordByUseId(userId int) *ReadRecord {

	var readRecord ReadRecord
	err := Mysql.DB.Where("user_id = ?", userId).First(&readRecord).Error
	if err != nil {
		fmt.Println(err)

		return nil

	}

	return &readRecord

}

func (this *ReadRecord) CreateRankRecordByUserId(userId int, day string, count int) {

}

func (this *ReadRecord) UpdateRankRecordByUserId(userId int, day string, count int) {

}
