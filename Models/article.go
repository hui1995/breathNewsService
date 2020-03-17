package Models

import (
	"breathNewsService/Databases/Mysql"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Channel    int
	Content    string
	Title      string
	Cover      string
	AuthorId   int    `gorm:"column:author_id"`
	AuthorName string `gorm:"column:author_name"`
	SourceType int    `gorm:"column:source_type"`
	Like       int
	Read       int
	Price      int
	State      int
}

func init() {
	table := Mysql.DB.HasTable(Article{})
	if !table {
		Mysql.DB.CreateTable(Article{})
	}
}

func (this *Article) FindById(id uint) (*Article, error) {
	var article Article
	err := Mysql.DB.Where("id = ?", id).First(&article).Error
	if err != nil {

		return nil, err

	}

	return &article, nil
}

func (this *Article) FindByChannel(channelId int) (artilelst []Article) {
	Db := Mysql.DB

	if channelId != 0 {
		Db = Db.Where("channel = ?", channelId)
	}
	Db = Db.Order("RAND()").Limit(8).Find(&artilelst)

	return artilelst

}
