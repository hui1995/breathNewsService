package Models

import (
	"breathNewsService/Databases/Mysql"
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

	UserName    string
	RealNum     int
	RealName    string
	Password    string
	Phone       string
	Alipay      string
	Money       float64
	FreezeMoney float64
	ActiveMoney float64
	Status      int
}

//func init() {
//	table := Mysql.DB.HasTable(User{})
//	if !table {
//		Mysql.DB.CreateTable(User{})
//	}
//}

func (this *User) FindByPhone(phone string) (*User, error) {

	var User User
	err := Mysql.DB.Where("phone = ?", phone).First(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil

}
func (this *User) FindLastUser() int {
	var user User
	Mysql.DB.Last(&user)

	return user.RealNum
}

func (this *User) InsertInfo(realNum int, status int, userName string) bool {
	user := User{UserName: userName, RealNum: realNum, Status: status}

	Mysql.DB.Create(&user)
	return true

}
