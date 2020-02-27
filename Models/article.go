package Models

import (
	"breathNewsService/Databases/Mysql"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Channel    int
	Content    string
	Title      string
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

func (this *Article) FindById(id int) (article Article, err error) {
	err = Mysql.DB.Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, err

	}
	fmt.Println(article.Title)

	return article, nil
}
