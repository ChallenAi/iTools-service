package controllers

import (
	// "fmt"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
)

func GetShares(ctx *fasthttp.RequestCtx) {
	shares, err := models.GetShares()
	if err != nil {
		utils.ServerFail(ctx)
	}
	utils.RespData(ctx, shares)
}

func AddShare(ctx *fasthttp.RequestCtx) {}

func DeleteShare(ctx *fasthttp.RequestCtx) {}

func EditShare(ctx *fasthttp.RequestCtx) {}