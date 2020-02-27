package Models

import "time"

/**
 * @Author: hui
 * @Description: 订单模型
 * @WebSite : https://www.breathcoder.cn
 * @File:  order
 * @Version: 1.0.0
 * @Date: 2020/2/25 12:43 PM
 */

 type Order struct {
 	Id int
 	UserId int
 	ProductId int
 	createTime time.Time
 	State int
 	UpdateTime time.Time
 	OptionUserId int

 }
