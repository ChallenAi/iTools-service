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
