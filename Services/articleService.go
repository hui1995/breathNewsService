package Services

import (
	"breathNewsService/Models"
	"breathNewsService/utils"
	"time"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/5 3:25 PM
 */

type articleHomeBean struct {
	Id         uint
	Title      string
	Cover      string
	AuthorName string
	Like       int
	Read       int
	Price      int
	IsRead     bool
	CreateTime time.Time
}

func GetArticleDeatil(id uint) (*Models.Article, error) {
	var article Models.Article

	return article.FindById(id)

}

func GetHomeArticleList(channel, userId int) []articleHomeBean {

	var article Models.Article
	var readConsumption Models.ReadConsumption
	//查询文章列表，假如没登录，则全部锁定，如果登录了
	var idList []interface{}

	articlelist := article.FindByChannel(channel)

	//获取查看过的userId
	if userId != 0 {
		readRecordList := readConsumption.FindReadRecordByUserId(userId)
		for _, v := range readRecordList {
			idList = append(idList, v.ArticleId)
		}

	}

	var articlelist2 []articleHomeBean
	for _, v := range articlelist {
		var articleHomebean1 articleHomeBean
		articleHomebean1.AuthorName = v.AuthorName
		articleHomebean1.Title = v.Title
		articleHomebean1.Read = v.Read
		articleHomebean1.Like = v.Like
		articleHomebean1.Price = v.Price
		articleHomebean1.CreateTime = v.CreatedAt
		articleHomebean1.Id = v.ID
		articleHomebean1.Cover = v.Cover
		if utils.Contarins(v.ID, idList) {
			articleHomebean1.IsRead = true
		} else {
			articleHomebean1.IsRead = false
		}
		articlelist2 = append(articlelist2, articleHomebean1)
	}

	return articlelist2

	//查询列表，关联查询

}
