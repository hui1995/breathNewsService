package Models

import (
	"breathNewsService/Databases/Mysql"
	"github.com/jinzhu/gorm"
)

//收入记录表

type IncomeRecord struct {
	gorm.Model
	Type    int
	Money   float64
	UserId  int
	Message string
	Status  int
}

//func init() {
//	table := Mysql.DB.HasTable(IncomeRecord{})
//	if !table {
//		Mysql.DB.CreateTable(IncomeRecord{})
//	}
//}

func (this *IncomeRecord) Insert(type1 int, userId int, money float64, stattus int, message string) bool {
	var incomeRecord = IncomeRecord{Type: type1, UserId: userId, Money: money, Status: stattus, Message: message}
	Mysql.DB.Create(&incomeRecord)
	return true

}

//获取记录
func (this *IncomeRecord) FindByUser(userId, status int) []IncomeRecord {
	incomeRecordlst := make([]IncomeRecord, 0)
	Db := Mysql.DB

	Db = Db.Where("user_id = ? and status = ?", userId, status).Order("id desc").Limit(30).Find(&incomeRecordlst)

	return incomeRecordlst
}

func (this *IncomeRecord) FindByUserType(userId int) IncomeRecord {
	var incomeRecord IncomeRecord
	Mysql.DB.Model(&IncomeRecord{}).Where("user_id = ? and type = 0 and status = 0", userId).Order("created_at desc").First(&incomeRecord)
	return incomeRecord

}

func (this *IncomeRecord) UpdateRecord(Id uint, status int) bool {
	var incomeRecord IncomeRecord
	Mysql.DB.First(&incomeRecord, Id).Update("status", status)
	return true
}
