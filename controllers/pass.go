package controllers

import (
	// "fmt"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
)

func GetPasswords(ctx *fasthttp.RequestCtx) {
	passwords, err := models.GetPasswords()
	if err != nil {
		utils.ServerFail(ctx)
	}
	utils.RespData(ctx, passwords)
}

func AddPass(ctx *fasthttp.RequestCtx) {}

func DeletePass(ctx *fasthttp.RequestCtx) {}

func EditPass(ctx *fasthttp.RequestCtx) {}