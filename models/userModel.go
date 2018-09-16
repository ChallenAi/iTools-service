package models

import (
	// "fmt"
	"time"
)

type User struct {
	UserId             int `gorm:"type:int(8);PRIMARY_KEY"`
	UserName           string
	Password           string
	PhoneNum           string
	Avatar             string
	CreateAt           time.Time `gorm:"Column:gmt_create"`
	UpdateAt           time.Time `gorm:"Column:gmt_modified"`
	IsValid            bool
	IsDeleted          bool `gorm:"Column:is_delete"`
	PasswordVerifyCode string
}

func (User) TableName() string {
	return "user"
}

// user.save
// func (user *User) save() error {
// 	return db.Model(user).Updates(map[string]interface{}{
// 		"view": post.View,
// 	}).Error
// }

func FindAllUsers() ([]User, error) {
	var users []User
	err := DB.Find(&users).Error
	// if err != nil {
	// 	fmt.Print(err)
	// }
	return users, err
}
