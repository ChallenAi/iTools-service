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
			"deleted": utils.RuleItem{Type: "numeric", Required: false},
			"page":    utils.RuleItem{Type: "numeric", Required: false},
			"perpage": utils.RuleItem{Type: "numeric", Required: false},
			"uid":     utils.RuleItem{Type: "numericId", Required: false},
			"typeId":  utils.RuleItem{Type: "numericId", Required: false},
			"keyword": utils.RuleItem{Type: "string", Required: false},
		},
	}

	data, errors := validator.Validate(ctx.QueryArgs())
	fmt.Println(data, errors)

	articles, err := models.SearchArticles()

	if err != nil {
		fmt.Println(err)
		utils.ServerFail(ctx)
	} else {
		utils.RespData(ctx, articles)
	}
}

func GetArticlesTitles(ctx *fasthttp.RequestCtx) {
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
}
