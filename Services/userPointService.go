package Services

import (
	"breathNewsService/Models"
	"fmt"
	"math/rand"
	"time"
)

func Consumption(id uint, userId int) (float64, int) {
	var article Models.Article
	var readPointsReward Models.ReadPointsReward
	var userPoints Models.UserPoints
	var readConsumption Models.ReadConsumption
	var incomeRecord Models.IncomeRecord

	isExist := readConsumption.ExistReadRecordByUserIdAndArticleId(userId, id)
	if isExist {
		return 0, -5
	}

	//查找文章价格
	articleInfo, err := article.FindById(id)
	if err != nil {
		return 0, -1

	}
	//查看该用户是否有足够的积分
	userPoints1 := userPoints.SelectPointsByUserId(userId)
	if userPoints1.Points < articleInfo.Price {
		return 0, -2
	}

	if userPoints.SubPoints(userId, articleInfo.Price) {
		readConsumption.InsertRecord(userId, id, articleInfo.Price)
	} else {
		return 0, -4
	}

	//获取积分
	readPointsReward1 := readPointsReward.FindByPoints(articleInfo.Price)
	high := readPointsReward1.Hight
	low := readPointsReward1.Low
	rate := readPointsReward1.Rate

	var price float64
	//如果在概率之内，则获取奖励
	if rand.Float64() < rate {
		for {
			price = rand.Float64()
			if low <= price && price < high {
				incomeRecord.Insert(0, userId, price, 0)
				break
			}
		}
		return price, 0

	} else {
		//如果不再则不获取奖励
		return 0, -3
	}

}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}
func AddPoints(userId int) bool {
	var userPoints Models.UserPoints
	var points float64
	//如果在概率之内，则获取奖励
	for {
		points = rand.Float64()
		if 0.6 <= points && points <= 0.8 {
			break
		}
	}

	fmt.Println(points)

	ppintsInt := int(points * 100)
	fmt.Println(ppintsInt)

	userPoints.AddPoints(userId, ppintsInt)
	return true

}
