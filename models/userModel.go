package models

type User struct {
	UserId             int
	UserName           string
	Password           string
	PhoneNum           string
	Avatar             string
	CreateAt           string
	UpdateAt           string
	IsValid            bool
	IsDeleted          bool
	PasswordVerifyCode string
}
