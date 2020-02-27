package Models

import (
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
	ProductId    int
	State        int
	OptionUserId int
}

//func init() {
//	table := Mysql.DB.HasTable(Order{})
//	if !table {
//		Mysql.DB.CreateTable(Order{})
//	}
//}
