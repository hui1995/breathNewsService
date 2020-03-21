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

func (this *User) FindByRealNum(realNum int) (*User, error) {

	var User User
	err := Mysql.DB.Where("real_num = ?", realNum).First(&User).Error
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

func (this *User) SelectByAliPay(alipay string) bool {
	var count int
	Mysql.DB.Model(&User{}).Where("alipay = ?", alipay).Count(&count)
	if count > 2 {

		return false

	}

	return true
}

// update aliipay
func (this *User) UpdateAlipay(userId int, alipay, realName string) bool {
	var user User
	Mysql.DB.Where(&User{RealNum: userId}).First(&user)

	Mysql.DB.Model(&user).Updates(User{Alipay: alipay, RealName: realName})
	return true
}
