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
 * @Date: 2020/2/27 2:28 PM
 */

type UserPoints struct {
	gorm.Model

	UserId int
	Points int
}

//func init() {
//	table := Mysql.DB.HasTable(UserPoints{})
//	if !table {
//		Mysql.DB.CreateTable(UserPoints{})
//	}
//}
