package Models

import (
	"breathNewsService/Databases/Mysql"
	"github.com/jinzhu/gorm"
	"time"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/2/26 3:59 PM
 */

type ReadConsumption struct {
	gorm.Model

	ArticleId uint
	UserId    int
	Point     int
	Type      int
	Status    int
}

//func init() {
// table := Mysql.DB.HasTable(ReadConsumption{})
// if !table {
//  Mysql.DB.CreateTable(ReadConsumption{})
// }
//}
func (this *ReadConsumption) InsertRecord(userId int, articleId uint, points int, Ttype int) bool {

	var readConsumption = ReadConsumption{ArticleId: articleId, UserId: userId, Point: points, Type: Ttype}
	Mysql.DB.Create(&readConsumption)
	return true

}

func (this *ReadConsumption) ExistReadRecordByUserIdAndArticleId(userId int, articleId uint) bool {

	Db := Mysql.DB
	var count = 0
	Db = Db.Model(&ReadConsumption{}).Where("user_id = ? AND article_id = ?", userId, articleId).Count(&count)
	if count == 0 {
		return false
	}

	return true

}

//获取消费记录
func (this *ReadConsumption) FindReadRecordByUserId(userId int) []ReadConsumption {

	readConsumptionlst := make([]ReadConsumption, 0)
	Db := Mysql.DB
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -7)
	Db = Db.Where("user_id = ? AND created_at >= ?", userId, oldTime).Find(&readConsumptionlst)
	return readConsumptionlst

}

func (this *ReadConsumption) UpdateRecord(articleId uint, userId int, status int) bool {
	var readConsumption ReadConsumption
	Mysql.DB.Model(&readConsumption).Where(&ReadConsumption{UserId: userId, ArticleId: articleId}).Update("status", status)
	return true
}
