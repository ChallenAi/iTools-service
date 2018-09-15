package controllers

import (
	// "fmt"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
)

func GetAllUsers(ctx *fasthttp.RequestCtx) {
	user := &models.User{
		UserId:   1,
		UserName: "Challen",
	}

	utils.RespJson(ctx, user)
}

func Login(ctx *fasthttp.RequestCtx) {
	postArgs := ctx.PostArgs()

	phoneNum := string(postArgs.Peek("phoneNum"))
	if phoneNum == "" {
		utils.ReqFail(ctx, "request error: phoneNum or password is wrong")
		return
	}

	password := string(postArgs.Peek("password"))
	if password == "" {
		utils.ReqFail(ctx, "request error: password or password is wrong")
		return
	}

	user := &models.User{
		UserId:   1,
		UserName: "有",
	}

	if phoneNum == "Test" {
		if password == "1234" {
			utils.RespJson(ctx, user)
			return
		} else {
			utils.ReqFail(ctx, "request error: password error")
			return
		}
	} else {
		utils.ReqFail(ctx, "request error: user not found")
	}
}

func Register(ctx *fasthttp.RequestCtx) {
	utils.ReqFail(ctx, "request error: register denny")

}
