package controllers

import (
	// "fmt"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
	"strconv"
)

func GetArticles(ctx *fasthttp.RequestCtx) {
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
		utils.ReqFail(ctx, "request error: article id must be sequence_id")
		return
	}

	article := &models.Article{
		ArticleId: articleId,
		Title:     "nice",
	}

	utils.RespJson(ctx, article)
}

func PostArticle(ctx *fasthttp.RequestCtx) {
	// postArgs := ctx.PostArgs()
}

func GetAllTags(ctx *fasthttp.RequestCtx) {
}
