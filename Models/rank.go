package Models

import "time"

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description: 
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/2/26 3:56 PM
 */

 type Rank struct {
 	Id int
 	UserId int
 	UserName string
 	CreateTime time.Time
 	IsRobot int
 	Score int

 }
