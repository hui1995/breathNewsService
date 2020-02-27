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
 * @Date: 2020/2/26 4:02 PM
 */

type ReadPointsReward struct {
	gorm.Model

	Points int
	Rate   float64
	Low    float64
	Hight  float64
}

//func init() {
//	table := Mysql.DB.HasTable(ReadPointsReward{})
//	if !table {
//		Mysql.DB.CreateTable(ReadPointsReward{})
//	}
//}
