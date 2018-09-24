package controllers

import (
	"fmt"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
)

func GetGists(ctx *fasthttp.RequestCtx) {
	validator := utils.Validator{
		Rules: map[string]utils.RuleItem{
			"deleted": utils.RuleItem{Type: "binary", Required: false, Alias: "is_deleted"},
			"uid":     utils.RuleItem{Type: "number", Required: false, Alias: "user_id"},
			"typeId":  utils.RuleItem{Type: "number", Required: false, Alias: "tag_id"},
			"page":    utils.RuleItem{Type: "pageNumber", Required: false},
			"perpage": utils.RuleItem{Type: "pageSize", Required: false},
			"keyword": utils.RuleItem{Type: "likeString", Required: false},
		},
	}

	data, errors := validator.Validate(ctx.QueryArgs())
	// fmt.Println(data, errors)
	if len(errors) > 0 {
		utils.RespFail(ctx, 400, errors[0])
		return
	}

	gists, err := models.GetGists(data)

	if err != nil {
		fmt.Println(err)
		utils.ServerFail(ctx)
	} else {
		utils.RespData(ctx, gists)
	}
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
