package models

type Article struct {
	ArticleId int
	Title     string
	Content   string
}

type ArticleCondition struct {
	Params  map[string]interface{}
	Keyword string
	Page    map[string]interface{}
}

func (Article) TableName() string {
	return "article"
}

func GetArticles(condition ArticleCondition) {

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
