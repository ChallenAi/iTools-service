package models

import (
	"time"
)

type Password struct {
	PasswordId int       `gorm:"type:int(8);PRIMARY_KEY" json:"articleId"`
	UserId     int       `gorm:"type:int(8)" json:"userId"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	TypeId     int       `json:"-"`
	View       int       `json:"view"`
	Like       int       `json:"like"`
	Collect    int       `json:"collect"`
	Rank       int       `json:"rank"`
	IsDeleted  bool      `gorm:"Column:deleted" json:"-"`
	CreateAt   time.Time `gorm:"Column:created_at" json:"createAt"`
	UpdateAt   time.Time `gorm:"Column:updated_at" json:"updateAt"`
}

func (Password) TableName() string {
	return "password"
}

func GetPasswords() ([]Password, error) {
	var passwords []Password
	err := DB.Find(&passwords).Error
	return passwords, err
}
