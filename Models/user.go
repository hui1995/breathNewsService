package Models

import (
	"github.com/jinzhu/gorm"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/2/26 4:29 PM
 */
type User struct {
	gorm.Model

	UserName string
	RealName string
	Password string
	Phone    string
	Alipay   string

	Money       float64
	FreezeMoney float64
	ActiveMoney float64
}

//func init() {
//	table := Mysql.DB.HasTable(User{})
//	if !table {
//		Mysql.DB.CreateTable(User{})
//	}
//}
