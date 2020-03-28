package Services

import (
	"breathNewsService/Models"
	"breathNewsService/utils"
	"encoding/json"
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
				incomeRecord.Insert(0, userId, price, 0, "阅读中奖")
				break
			}
		}
		return price, 0

	} else {
		//如果不再则不获取奖励
		return 0, -3
	}

}

func AddPoints(userId int) int {
	var userPoints Models.UserPoints
	var points float64
	var pointsRecord Models.PointsRecord
	day := utils.TransDayInt(time.Now().Year(), int(time.Now().Month()), time.Now().Day()-1)

	pointsInfo := pointsRecord.InsertPointsRAecord(userId, day)
	count := pointsInfo.Count + 1
	pointsRecord.UpdatePointsRecord(userId, day, count)

	//填充次数

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
	return ppintsInt

}

type QueryRule struct {
	Limit    int
	Interval int
}

func SelectLimit(userId int) (int, int) {
	var pointsRecord Models.PointsRecord
	var commconfig2 Models.CommonConfig
	var limit int
	day := utils.TransDayInt(time.Now().Year(), int(time.Now().Month()), time.Now().Day()-1)

	pointInfo := pointsRecord.InsertPointsRAecord(userId, day)
	value := commconfig2.FindByGroupAndKey("ad", "queryRule")
	var queryRule []QueryRule
	err3 := json.Unmarshal([]byte(value), &queryRule)
	if err3 != nil {
		limit = 0

	} else {
		limit = queryRule[0].Limit
	}
	result := limit - pointInfo.Count
	if result < 0 {
		return 0, 0
	}

	Interval := queryRule[0].Interval
	now := time.Now()

	subM := now.Sub(pointInfo.UpdatedAt)
	interval2 := subM.Seconds()
	if int(interval2) < (Interval) {
		return -1, Interval - int(interval2)

	}
	return result, 0
}

func ReadIsReal(userId int, articleId int) {
	var readConsumption Models.ReadConsumption
	readConsumption.UpdateRecord(uint(articleId), userId, 1)
}
