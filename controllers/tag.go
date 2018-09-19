package controllers

import (
	// "fmt"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
)

func GetTags(ctx *fasthttp.RequestCtx) {
	tags, err := models.GetTags()
	if err != nil {
		utils.ServerFail(ctx)
	}
	utils.RespData(ctx, tags)
}