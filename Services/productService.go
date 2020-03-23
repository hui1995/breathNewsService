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
	Tip         string
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
	result.Tip = "注意事项\n 1.提现申请将在1到3个工作日内审批到账，请耐心等待 \n 2.我的订单 可查看兑换记录状态 \n 3.目前技术受限，暂时只支持支付宝提现，后期将开通微信"

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

type OrderResult struct {
	Price      float64
	StateMsg   string
	AppleyTime string
}

func GetOrderList(userId int) []OrderResult {
	var order Models.Order

	orderlst := order.FindeOrderByUserId(userId)
	var orderRestultlst []OrderResult
	for _, v := range orderlst {
		var orderResult OrderResult
		if v.State == 0 {
			orderResult.StateMsg = "待审核"
		} else if v.State == 1 {
			orderResult.StateMsg = "已通过"

		} else {
			orderResult.StateMsg = "已拒绝"
		}

		orderResult.Price = v.Price
		applayTime := v.CreatedAt.Format("2006年01月02日")
		orderResult.AppleyTime = applayTime
		orderRestultlst = append(orderRestultlst, orderResult)

	}
	return orderRestultlst

}
