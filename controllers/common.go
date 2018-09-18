package controllers

import (
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
)

func Pong(ctx *fasthttp.RequestCtx) {
	utils.RespData(ctx, "pong")
}

func NotFound(ctx *fasthttp.RequestCtx) {
	utils.RespFail(ctx, 404, "没有找到该接口.")
}
