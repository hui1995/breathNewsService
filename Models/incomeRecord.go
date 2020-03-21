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
}

//func init() {
//	table := Mysql.DB.HasTable(IncomeRecord{})
//	if !table {
//		Mysql.DB.CreateTable(IncomeRecord{})
//	}
//}

func (this *IncomeRecord) Insert(type1 int, userId int, money float64) bool {
	var incomeRecord = IncomeRecord{Type: type1, UserId: userId, Money: money}
	Mysql.DB.Create(&incomeRecord)
	return true

}

//获取记录
func (this *IncomeRecord) FindByUser(userId int) []IncomeRecord {
	incomeRecordlst := make([]IncomeRecord, 0)
	Db := Mysql.DB

	Db = Db.Where("user_id = ?", userId).Order("id desc").Limit(30).Find(&incomeRecordlst)

	return incomeRecordlst
}
