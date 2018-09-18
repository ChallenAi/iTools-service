package models

import (
	"time"
)

type Gist struct {
	GistId    int       `gorm:"type:int(8);PRIMARY_KEY" json:"articleId"`
	UserId    int       `gorm:"type:int(8)" json:"userId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	TagId     int       `json:"-"`
	IsDeleted bool      `gorm:"Column:deleted" json:"-"`
	CreateAt  time.Time `gorm:"Column:gmt_create" json:"createAt"`
	UpdateAt  time.Time `gorm:"Column:gmt_modified" json:"updateAt"`
}

// type GistCondition struct {
// 	Params  map[string]interface{}
// 	Keyword string
// 	Page    map[string]interface{}
// }

func (Gist) TableName() string {
	return "tip"
}

func GetGists() ([]Gist, error) {
	var gists []Gist
	err := DB.Find(&gists).Error
	return gists, err
}
