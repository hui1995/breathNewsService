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
	lst := incomeRcord.FindByUser(userId)
	var resultlst []IncomeHisty
	for _, v := range lst {
		var incomeHistory IncomeHisty
		incomeHistory.Money = v.Money
		incomeHistory.Message = v.Message
		resultlst = append(resultlst, incomeHistory)

	}
	return resultlst
}
