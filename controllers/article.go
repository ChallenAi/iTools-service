package controllers

import (
	// "fmt"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
	"strconv"
)

func GetArticles(ctx *fasthttp.RequestCtx) {

	q := ctx.QueryArgs()

	condition := &models.ArticleCondition{}
	condition.Params = map[string]interface{}{}

	if q.Peek("deleted") != nil && string(q.Peek("deleted")) == "1" {
		condition.Params["IsDelete"] = true
	} else {
		condition.Params["IsDelete"] = false
	}

	// if q.Peek("page") != nil {
	// 	condition.Params.UserId = q.uid
	// }

	// if q.Peek("perpage") != nil {
	// 	condition.Params.UserId = q.uid
	// }

	if p3 := q.Peek("uid"); p3 != nil {
		condition.Params["IsDelete"] = string(p3)
	}

	if p4 := q.Peek("keyword"); p4 != nil {
		condition.Keyword = string(p4)
	}

	if p5 := q.Peek("typeId"); p5 != nil {
		condition.Params["typeId"] = string(p5)
	}

	articles, err := models.GetArticles(condition)

	if err != nil {
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
