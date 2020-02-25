package Models




type Test struct {
	Id int `gorm:"column:id"`
	UserId int `gorm:"column:user_id"`
	Points int `gorm:"column:points"`
}
