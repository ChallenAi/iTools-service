package models

import (
	// "fmt"
	"time"
)

type User struct {
	UserId             int       `gorm:"type:int(8);PRIMARY_KEY" json:"userId"`
	UserName           string    `json:"userName"`
	Password           string    `json:"password"`
	PhoneNum           string    `json:"phoneNum"`
	Avatar             string    `json:"avatar"`
	CreateAt           time.Time `gorm:"Column:gmt_create" json:"createAt"`
	UpdateAt           time.Time `gorm:"Column:gmt_modified" json:"updateAt"`
	IsValid            bool      `json:"-"`
	IsDeleted          bool      `gorm:"Column:is_delete" json:"-"`
	PasswordVerifyCode string    `json:"-"`
}

func (User) TableName() string {
	return "user"
}

func (user *User) Persist() error {
	DB.NewRecord(*user)
	return DB.Create(user).Error
}

func FindAllUsers() ([]User, error) {
	var users []User
	err := DB.Find(&users).Error
	return users, err
}
