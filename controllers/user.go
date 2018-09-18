package controllers

import (
	"time"
	// "fmt"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
	"github.com/lib/pq"
)

func GetAllUsers(ctx *fasthttp.RequestCtx) {
	users, err := models.FindAllUsers()
	if err != nil {
		utils.ServerFail(ctx)
	}
	utils.RespData(ctx, users)
}

func Login(ctx *fasthttp.RequestCtx) {
	postArgs := ctx.PostArgs()

	phoneNum := string(postArgs.Peek("phoneNum"))
	if phoneNum == "" {
		utils.RespFail(ctx, 400, "request error: phoneNum or password is wrong")
		return
	}

	password := string(postArgs.Peek("password"))
	if password == "" {
		utils.RespFail(ctx, 400, "request error: password or password is wrong")
		return
	}

	user := &models.User{
		UserId:   1,
		UserName: "有",
	}

	if phoneNum == "Test" {
		if password == "1234" {
			utils.RespData(ctx, user)
			return
		} else {
			utils.RespFail(ctx, 400, "request error: password error")
			return
		}
	} else {
		utils.RespFail(ctx, 400, "request error: user not found")
	}
}

func Register(ctx *fasthttp.RequestCtx) {
	user := &models.User{
		UserName: "小子",
		Password: "1234",
		PhoneNum: "15317752613",
		Avatar: "/img/timg.jpeg",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		IsValid: true,
		IsDeleted: false,
		PasswordVerifyCode: "abcd",
	}

	err := user.Persist()
	if err != nil {
		if err.(*pq.Error).Code.Name() == "unique_violation" {
			utils.RespFail(ctx, 400, "duplicated: phoneNum")
		} else {
			utils.ServerFail(ctx)
		}
		return
	}
	utils.RespSucc(ctx)
	// utils.RespFail(ctx, 400, "request error: register denny")
}
