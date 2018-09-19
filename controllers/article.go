package controllers

import (
	// "fmt"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
	"strconv"
)

func GetArticles(ctx *fasthttp.RequestCtx) {

	// const q = req.query
	// // 查询参数
	// const condition = {}
	// // 查询参数：带默认值的，必须有的
	// condition.deleted = (q.deleted === '1')?true:false
	// const page = parseInt(q.page || 1)
	// const limit = parseInt(q.perpage || 10)
	// const offset = limit * (page - 1)
	// // 如果查询特定用户/类型的文章
	// if (q.uid) condition.userId = q.uid
	// if (q.typeId) condition.typeId = q.typeId
	// if (q.keyword) condition.keyword = q.keywor

	validator := utils.Validator{
		Rules: map[string]utils.RuleItem{
			"deleted": utils.RuleItem{ Type: "int", Required: true },
			"page": utils.RuleItem{ Type: "int", Required: false },
		},
	}
	validator.Validate(ctx.QueryArgs())




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
