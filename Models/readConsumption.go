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
 * @Date: 2020/2/26 3:59 PM
 */

type ReadConsumption struct {
	gorm.Model

	ArticleId int
	UserId    int
	Point     int
}

//func init() {
// table := Mysql.DB.HasTable(ReadConsumption{})
// if !table {
//  Mysql.DB.CreateTable(ReadConsumption{})
// }
//}
