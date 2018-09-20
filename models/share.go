package models

import (
	"time"
)

type Share struct {
	ShareId   int       `gorm:"type:int(8);PRIMARY_KEY" json:"articleId"`
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

func (Share) TableName() string {
	return "share"
}

func GetShares() ([]Share, error) {
	var shares []Share
	err := DB.Find(&shares).Error
	return shares, err
}
