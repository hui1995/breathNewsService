package Models

import "time"

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description: 
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/2/26 4:29 PM
 */
 type User struct {
 	Id int
 	UserName string
 	RealName string
 	Password string
 	Phone string
 	Alipay string
 	CreateTime time.Time
 	Money float64
 	FreezeMoney float64
 	ActiveMoney float64
 }
