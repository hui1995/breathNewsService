package Models

import "time"



type Article struct {
	Id int
	Channel int
	Content string
	AuthorId int
	AuthorName string
	CreateTime time.Time
	SourceType int
	Like int
	Read int
	Price int
	State int

}