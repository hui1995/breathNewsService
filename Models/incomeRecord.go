package Models

import "time"


//收入记录表

type IncomeRecord struct {
	Id int
	Type int
	Money float64
	CreateTime time.Time
	userId int
}
