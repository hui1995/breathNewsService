package Models

import "time"

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description: 
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/2/26 3:59 PM
 */

 type ReadConsumption struct {
  Id int
  ArticleId int
  UserId int
  Point int
  CreateTime time.Time
 }