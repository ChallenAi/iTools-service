package models

import (
	"time"
)

type Article struct {
	ArticleId int
	UserId    int `gorm:"type:int(8);PRIMARY_KEY"`
	Title     string
	Content   string
	TypeId    int
	View      int
	Like      int
	Collect   int
	Rank      int
	IsDeleted bool      `gorm:"Column:deleted"`
	CreateAt  time.Time `gorm:"Column:created_at"`
	UpdateAt  time.Time `gorm:"Column:updated_at"`
}

type ArticleCondition struct {
	Params  map[string]interface{}
	Keyword string
	Page    map[string]interface{}
}

func (Article) TableName() string {
	return "article"
}

func GetArticles(condition *ArticleCondition) ([]Article, error) {
	var articles []Article
	err := DB.Find(&articles).Error
	// if err != nil {
	// 	fmt.Print(err)
	// }
	return articles, err
}

func (article *Article) addArticle() {

}

func getTitles() {

}

func getAllTags() {

}

func getArticle() {

}

func articleViewPlusOne() {

}
