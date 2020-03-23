package Services

import (
	"breathNewsService/Models"
	"encoding/json"
	"fmt"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/21 8:47 PM
 */

type ResultProduct struct {
	Money       float64
	FreezeMoney float64
	ActiveMoney float64
	Alipay      string
	RealName    string
	Pricelist   []int
}
type ProductInfo struct {
	Product int
	Price   int
}

func FindProductInfo(userId int) ResultProduct {
	var user Models.User
	var config Models.CommonConfig
	userInfo, err := user.FindByRealNum(userId)
	var result ResultProduct
	result.Money = userInfo.Money
	result.FreezeMoney = userInfo.FreezeMoney
	result.ActiveMoney = userInfo.ActiveMoney
	result.Alipay = userInfo.Alipay
	result.RealName = userInfo.RealName

	value := config.FindByGroupAndKey("product", "price")
	var productInfo []ProductInfo
	err2 := json.Unmarshal([]byte(value), &productInfo)
	if err2 != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)

	}
	var pricelst []int

	for _, v := range productInfo {
		pricelst = append(pricelst, v.Price)

	}
	result.Pricelist = pricelst

	return result

}

func AddOrderInfo(userId int, price float64) bool {
	var order Models.Order
	var user Models.User
	userInfo, _ := user.FindByRealNum(userId)
	freezeMoney := userInfo.FreezeMoney + price
	activeMoney := userInfo.ActiveMoney - price

	if order.FindeOrderByState(userId, 0) {
		user.UpdateMoney(userId, userInfo.Money, freezeMoney, activeMoney)
		return order.InsertOrder(userId, price)

	} else {
		return false
	}
}
