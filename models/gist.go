package models

import (
	"github.com/ChallenAi/iTools-service/utils"
	"time"
)

type Gist struct {
	GistId    int       `gorm:"type:int(8);PRIMARY_KEY" json:"gistId"`
	UserId    int       `gorm:"type:int(8)" json:"userId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	TagId     int       `json:"-"`
	IsDeleted bool      `gorm:"Column:deleted" json:"-"`
	CreateAt  time.Time `gorm:"Column:gmt_create" json:"createAt"`
	UpdateAt  time.Time `gorm:"Column:gmt_modified" json:"updateAt"`
	UserName  string    `gorm:"-" json:"username"`
	TypeName  string    `gorm:"-" json:"type"`
}

// type GistCondition struct {
// 	Params  map[string]interface{}
// 	Keyword string
// 	Page    map[string]interface{}
// }

func (Gist) TableName() string {
	return "tip"
}

func GetGists(serviceParams *utils.ServiceParams) ([]Gist, error) {
	gists := []Gist{}
	err := DB.
		Where(serviceParams.CommonParams).
		Select(`"user".user_name, tip.*, tag.content as type_name`).
		Joins("JOIN tag on tip.tag_id = tag.tag_id").
		Joins("JOIN \"user\" on tip.user_id = \"user\".user_id").
		Find(&gists).Error

	return gists, err
}

func AddGist(gist *Gist) bool {
	ok := DB.NewRecord(*gist)
	// db.Create(&user)
	return ok
}
