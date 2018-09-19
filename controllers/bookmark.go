package controllers

import (
	// "fmt"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
)

func GetBookmarks(ctx *fasthttp.RequestCtx) {
	bookmarks, err := models.GetBookmarks()
	if err != nil {
		utils.ServerFail(ctx)
	}
	utils.RespData(ctx, bookmarks)
}

func AddBookmark(ctx *fasthttp.RequestCtx) {}

func DeleteBookmark(ctx *fasthttp.RequestCtx) {}

func EditBookmark(ctx *fasthttp.RequestCtx) {}