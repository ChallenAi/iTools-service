package utils

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type ErrorMsg struct {
	Code int
	Msg  string
}

func ReqFail(ctx *fasthttp.RequestCtx, msg string) {
	respMsg := &ErrorMsg{
		Code: 400,
		Msg:  msg,
	}
	resp, _ := json.Marshal(respMsg)
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.Write(resp)
}

func ServerFail(ctx *fasthttp.RequestCtx) {
	respMsg := &ErrorMsg{
		Code: 500,
		Msg:  "server error!",
	}
	resp, _ := json.Marshal(respMsg)
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.Write(resp)
}

func RespJson(ctx *fasthttp.RequestCtx, v interface{}) {
	succMsg := map[string]interface{}{
		"code": 0,
		"data": v,
	}
	resp, err := json.Marshal(succMsg)
	if err != nil {
		ServerFail(ctx)
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.Write(resp)
}
