package Services

import (
	"breathNewsService/Models"
	"math/rand"
	"time"
)

func Consumption(id uint, userId int) (float64, int) {
	var article Models.Article
	var readPointsReward Models.ReadPointsReward
	var userPoints Models.UserPoints
	var readConsumption Models.ReadConsumption

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
