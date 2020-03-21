package Models

import "github.com/jinzhu/gorm"

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/21 10:16 PM
 */
type Invite struct {
	gorm.Model
	Inviter      int
	invitee      int
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
