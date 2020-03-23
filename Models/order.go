package Models

import (
	"breathNewsService/Databases/Mysql"
	"github.com/jinzhu/gorm"
)

/**
 * @Author: hui
 * @Description: 订单模型
 * @WebSite : https://www.breathcoder.cn
 * @File:  order
 * @Version: 1.0.0
 * @Date: 2020/2/25 12:43 PM
 */

type Order struct {
	gorm.Model
	UserId       int
	Price        float64
	State        int
	OptionUserId int
	Desction     string
}

//func init() {
//	table := Mysql.DB.HasTable(Order{})
//	if !table {
//		Mysql.DB.CreateTable(Order{})
//	}
//}

// update aliipay
func (this *Order) InsertOrder(userId int, Price float64) bool {
	order := Order{UserId: userId, Price: Price, State: 0, Desction: "待审核"}

	Mysql.DB.Create(&order)
	return true

}

func (this *Order) FindeOrderByState(userId, state int) bool {

	var count int
	Mysql.DB.Model(&Order{}).Where("user_id = ? and state = ?", userId, state).Count(&count)
	if count == 0 {

		return true

	}

	return false

}
