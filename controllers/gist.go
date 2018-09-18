package controllers

import (
	// "fmt"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
)

func GetGists(ctx *fasthttp.RequestCtx) {
	gists, err := models.GetGists()
	if err != nil {
		utils.ServerFail(ctx)
	}
	utils.RespData(ctx, gists)
}

func AddGist(ctx *fasthttp.RequestCtx) {
	// postArgs := ctx.PostArgs()

	// gist := &models.Gist{
	// 	Title: string(postArgs.Peek("title"))
	// 	Content: string(postArgs.Peek("content"))
	// 	UserId: string(postArgs.Peek("authorId"))
	// 	TagId: string(postArgs.Peek("typeId"))
	// }

	// phoneNum := string(postArgs.Peek("phoneNum"))
	// if phoneNum == "" {
	// 	utils.RespFail(ctx, 400, "request error: phoneNum or password is wrong")
	// 	return
	// }

	// password := string(postArgs.Peek("password"))
	// if password == "" {
	// 	utils.RespFail(ctx, 400, "request error: password or password is wrong")
	// 	return
	// }

	// ok := models.AddGist(gist)
	// if !ok {
	// 	utils.ServerFail(ctx)
	// }
	utils.RespSucc(ctx)
}

func DeleteGist(ctx *fasthttp.RequestCtx) {}

func EditGist(ctx *fasthttp.RequestCtx) {}
