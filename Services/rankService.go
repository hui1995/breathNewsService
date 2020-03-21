package Services

import (
	"breathNewsService/Models"
	"breathNewsService/utils"
	"strconv"
	"time"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/17 11:10 PM
 */

type TodayInfo struct {
	UserName string
	Value    string
	Position string
	IsMe     bool
}

func GetTodayRank(userId int) []TodayInfo {

	var rank Models.Rank
	var todayInfos []TodayInfo

	ranklist := rank.GetRank()
	for _, v := range ranklist {
		var rankInfo TodayInfo

		rankInfo.Value = strconv.Itoa(v.Score)
		if v.UserId == userId {
			rankInfo.IsMe = true

		} else {
			rankInfo.IsMe = false
		}
		rankInfo.Position = strconv.Itoa(v.Position)
		rankInfo.UserName = v.UserName
		todayInfos = append(todayInfos, rankInfo)

	}

	if userId == 0 {
		return todayInfos
	}
	myRank := rank.GetCurrentRank(userId)
	if myRank.Position > 10 {
		var todayInfo TodayInfo
		todayInfo.UserName = myRank.UserName
		todayInfo.Value = strconv.Itoa(myRank.Score)
		if myRank.Position >= 1888 {
			todayInfo.Position = "未上榜"
		} else {
			todayInfo.Position = strconv.Itoa(rank.Position)
		}
		todayInfo.IsMe = true

		todayInfos = append(todayInfos, todayInfo)

	}

	return todayInfos
}

func GetYesRank(userId int) []TodayInfo {
	var rankRecord Models.RankRecord

	day := utils.TransDayInt(time.Now().Year(), int(time.Now().Month()), time.Now().Day()-1)
	recordlst := rankRecord.GetYesRank(day)
	myyesRank := rankRecord.GetCurrentRank(day, userId)

	var rankYess []TodayInfo
	for _, v := range recordlst {
		var rankInfo2 TodayInfo
		if v.UserId == userId {
			rankInfo2.IsMe = true

		} else {
			rankInfo2.IsMe = false
		}
		rankInfo2.Position = strconv.Itoa(v.Position)

		rankInfo2.UserName = v.UserName
		rankInfo2.Value = strconv.FormatFloat(v.Reward, 'g', 1, 64)
		rankYess = append(rankYess, rankInfo2)

	}

	if userId == 0 {
		return rankYess
	}
	if myyesRank.Position > 10 {
		var rankYes TodayInfo
		if myyesRank.Position >= 1000 {
			rankYes.Position = "未上榜"

		} else {
			rankYes.Position = strconv.Itoa(myyesRank.Position)

		}
		rankYes.UserName = myyesRank.UserName
		rankYes.Value = strconv.FormatFloat(myyesRank.Reward, 'g', 1, 64)
		rankYess = append(rankYess, rankYes)

	}

	return rankYess

}
