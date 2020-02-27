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

func (this *Article) FindById(id int) (*Article, error) {
	var article Article
	err := Mysql.DB.Where("id = ?", id).First(&article).Error
	if err != nil {

		return nil, err

	}

	return &article, nil
}


func(this *Article) FindByChannel(channelId,page,PageSize int)([]Article){
	artilelst := make([]Article, 0)
	Db:=Mysql.DB
	if channelId!=0{
		Db=Db.Where("channel = ?",channelId)
	}
	Db=Db.Limit(PageSize).Offset((page-1)*PageSize).Find(&artilelst)

	return artilelst

}
