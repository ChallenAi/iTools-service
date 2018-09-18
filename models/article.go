package models

import (
	"time"
)

type Article struct {
	ArticleId int       `gorm:"type:int(8);PRIMARY_KEY" json:"articleId"`
	UserId    int       `gorm:"type:int(8)" json:"userId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	TypeId    int       `json:"-"`
	View      int       `json:"view"`
	Like      int       `json:"like"`
	Collect   int       `json:"collect"`
	Rank      int       `json:"rank"`
	IsDeleted bool      `gorm:"Column:deleted" json:"-"`
	CreateAt  time.Time `gorm:"Column:created_at" json:"createAt"`
	UpdateAt  time.Time `gorm:"Column:updated_at" json:"updateAt"`
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
