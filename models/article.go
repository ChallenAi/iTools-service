package models

import (
	"fmt"
	"time"
)

type Article struct {
	ArticleId int       `gorm:"type:int(8);PRIMARY_KEY" json:"articleId"`
	UserId    int       `json:"userId"`
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
	// User      User      `gorm:"ForeignKey:UserId" json:"-"`
	// Type      Tag       `gorm:"ForeignKey:TagId" json:"-"`
	UserName string `gorm:"-" json:"username"`
	TypeName string `gorm:"-" json:"type"`
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

	// raw: sql is long, but useful
	// var articles []Article
	// rows, err := DB.Table("article").Select("article_id, user_id").Rows()
	// rows, err := DB.Raw("select a.article_id, a.user_id, u.user_name from article a, \"user\" u where a.article_id = ? and u.user_id = a.user_id", "2").Rows()
	// defer rows.Close()

	// var (
	// 	articleId string
	// 	userId    string
	// 	username  string
	// )

	// for rows.Next() {
	// 	fmt.Println(rows.Columns())
	// 	rows.Scan(&articleId, &userId, &username)
	// 	fmt.Println(articleId, userId, username)
	// }

	// 2 use inner relation, preload use 3 querys....
	// articles := []Article{}
	// err := DB.Preload("User").Preload("Type").Find(&articles).Error

	// for idx, i := range articles {
	// 	fmt.Println(i.User)
	// 	articles[idx].UserName = i.User.UserName
	// 	articles[idx].TypeName = i.Type.Content
	// }

	// build sql
	articles := []Article{}
	err := DB.
		Select(`"user".user_name, article.*, tag.content as type_name`).
		Joins("JOIN tag on article.type_id = tag.tag_id").
		Joins("JOIN \"user\" on article.user_id = \"user\".user_id").
		Find(&articles).Error

	fmt.Println("")
	// fmt.Println(articles)

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
