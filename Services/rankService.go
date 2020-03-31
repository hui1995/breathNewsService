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
	Day      int
	Day2     int
	UserId   int
}

func GetTodayRank(userId int) ([]TodayInfo, TodayInfo) {

	var rank Models.Rank
	var todayInfos []TodayInfo

	var readRecord Models.ReadRecord

	ranklist := rank.GetRank()
	for _, v := range ranklist {
		var rankInfo TodayInfo

		rankInfo.Value = strconv.Itoa(v.Score)

		rankInfo.Position = strconv.Itoa(v.Position)
		rankInfo.UserName = v.UserName
		todayInfos = append(todayInfos, rankInfo)

	}
	var todayInfo TodayInfo

	if userId == 0 {
		return todayInfos, todayInfo
	}

	readDayInfo := readRecord.GetDatRecordByUseId(userId)

	if readDayInfo.Count >= 7 {
		myRank := rank.GetCurrentRank(userId)
		todayInfo.UserName = myRank.UserName
		todayInfo.Value = strconv.Itoa(myRank.Score)
		todayInfo.Position = strconv.Itoa(rank.Position)
		todayInfo.IsMe = true

	} else {
		todayInfo.IsMe = false
		todayInfo.Day = readDayInfo.Count
		todayInfo.UserId = readDayInfo.UserId
		todayInfo.Day2 = 7 - readDayInfo.Count
	}

	return todayInfos, todayInfo
}

func GetYesRank(userId int) ([]TodayInfo, TodayInfo) {
	var rankRecord Models.RankRecord
	var readRecord Models.ReadRecord

	day := utils.TransDayInt(time.Now().Year(), int(time.Now().Month()), time.Now().Day()-1)
	recordlst := rankRecord.GetYesRank(day)

	var rankYess []TodayInfo
	for _, v := range recordlst {
		var rankInfo2 TodayInfo

		rankInfo2.Position = strconv.Itoa(v.Position)

		rankInfo2.UserName = v.UserName
		//n2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", v.Reward), 64)
		rankInfo2.Value = strconv.FormatFloat(float64(v.Reward), 'f', 2, 64)
		rankYess = append(rankYess, rankInfo2)

	}
	var todayInfo TodayInfo

	if userId == 0 {
		return rankYess, todayInfo
	}
	readDayInfo := readRecord.GetDatRecordByUseId(userId)

	if readDayInfo.Count >= 7 {
		myyesRank := rankRecord.GetCurrentRank(day, userId)
		todayInfo.UserName = myyesRank.UserName
		todayInfo.Value = strconv.Itoa(myyesRank.Score)
		todayInfo.Position = strconv.Itoa(myyesRank.Position)
		todayInfo.IsMe = true

	} else {
		todayInfo.IsMe = false
		todayInfo.Day = readDayInfo.Count
		todayInfo.UserId = readDayInfo.UserId
		todayInfo.Day2 = 7 - readDayInfo.Count
	}

	return rankYess, todayInfo

}
