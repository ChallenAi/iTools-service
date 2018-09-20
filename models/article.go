package models

import (
	"fmt"
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

func (Article) TableName() string {
	return "article"
}

func GetArticles() ([]Article, error) {
	var articles []Article
	err := DB.Find(&articles).Error
	return articles, err
}

func SearchArticles() ([]Article, error) {

	// raw query test
	// var articles []Article
	// rows, err := DB.Table("article").Select("article_id, user_id").Rows()
	rows, err := DB.Raw("select a.article_id, a.user_id, u.user_name from article a, \"user\" u where a.article_id = ? and u.user_id = a.user_id", "2").Rows()
	defer rows.Close()

	var (
		articleId string
		userId    string
		username  string
	)

	for rows.Next() {
		fmt.Println(rows.Columns())
		rows.Scan(&articleId, &userId, &username)
		fmt.Println(articleId, userId, username)
	}
	return []Article{}, err
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
