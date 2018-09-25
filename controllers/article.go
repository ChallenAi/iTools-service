package controllers

import (
	"fmt"
	"strconv"

	"github.com/ChallenAi/iTools-service/models"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
)

func GetArticles(ctx *fasthttp.RequestCtx) {

	validator := utils.Validator{
		Rules: map[string]utils.RuleItem{
			"deleted": utils.RuleItem{Type: "binary", Required: false, Alias: "deleted"},
			"uid":     utils.RuleItem{Type: "number", Required: false, Alias: "user_id"},
			"typeId":  utils.RuleItem{Type: "number", Required: false, Alias: "type_id"},
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

	articles, err := models.SearchArticles(data)

	if err != nil {
		fmt.Println(err)
		utils.ServerFail(ctx)
	} else {
		utils.RespData(ctx, articles)
	}
}

func GetArticlesTitles(ctx *fasthttp.RequestCtx) {
	validator := utils.Validator{
		Rules: map[string]utils.RuleItem{
			// "deleted": utils.RuleItem{Type: "binary", Required: false, Alias: "deleted"},
			"typeId": utils.RuleItem{Type: "number", Required: false, Alias: "type_id"},
		},
	}

	data, errors := validator.Validate(ctx.QueryArgs())
	if len(errors) > 0 {
		utils.RespFail(ctx, 400, errors[0])
		return
	}

	articles, err := models.GetTitles(data)

	if err != nil {
		fmt.Println(err)
		utils.ServerFail(ctx)
	} else {
		utils.RespData(ctx, articles)
	}
}

func GetArticle(ctx *fasthttp.RequestCtx) {
	// articleIdString, ok := ctx.UserValue("id").(string)
	// if !ok {
	// 	fmt.Println("err")
	// }
	articleIdString, _ := ctx.UserValue("id").(string)

	articleId, err := strconv.Atoi(articleIdString)
	if err != nil {
		utils.RespFail(ctx, 400, "request error: article id must be sequence_id")
		return
	}

	article := &models.Article{
		ArticleId: articleId,
		Title:     "nice",
	}

	utils.RespData(ctx, article)
}

func PostArticle(ctx *fasthttp.RequestCtx) {
	// postArgs := ctx.PostArgs()
}

func GetAllTags(ctx *fasthttp.RequestCtx) {
	resu, err := models.GetAllTags()

	if err != nil {
		fmt.Println(err)
		utils.ServerFail(ctx)
	} else {
		utils.RespData(ctx, resu)
	}
}
