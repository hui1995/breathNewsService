package Services

import "breathNewsService/Models"

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/18 12:10 AM
 */

type IncomeHisty struct {
	Money   float64
	Message string
}

func GetIncomeRecord(userId int) []IncomeHisty {
	var incomeRcord Models.IncomeRecord
	lst := incomeRcord.FindByUser(userId, 1)
	var resultlst []IncomeHisty
	for _, v := range lst {
		var incomeHistory IncomeHisty
		incomeHistory.Money = v.Money
		incomeHistory.Message = v.Message
		resultlst = append(resultlst, incomeHistory)

	}
	return resultlst
}

func AddIncome(userId int) bool {
	var incomeRecord Models.IncomeRecord
	incomeInfo := incomeRecord.FindByUserType(userId)
	if incomeInfo.Money != 0 {
		incomeRecord.UpdateRecord(incomeInfo.ID, 1)
		var userinfo Models.User
		username, _ := userinfo.FindByRealNum(userId)
		activtyMoney := username.ActiveMoney + incomeInfo.Money
		Money := username.Money + incomeInfo.Money
		userinfo.UpdateMoney(userId, Money, username.FreezeMoney, activtyMoney)
	}

	return true

}
