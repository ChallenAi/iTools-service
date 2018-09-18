package utils

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type ErrorMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DataMsg struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ListDataMsg struct {
	Code  int         `json:"code"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

func RespFail(ctx *fasthttp.RequestCtx, code int, msg string) {
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
		Msg:  "服务器错误!",
	}
	resp, _ := json.Marshal(respMsg)
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.Write(resp)
}

func RespData(ctx *fasthttp.RequestCtx, v interface{}) {
	respMsg := &DataMsg{
		Code: 0,
		Data: v,
	}
	resp, err := json.Marshal(respMsg)
	if err != nil {
		ServerFail(ctx)
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.Write(resp)
}

func RespListData(ctx *fasthttp.RequestCtx, v interface{}, c int) {
	respMsg := &ListDataMsg{
		Code:  0,
		Count: c,
		Data:  v,
	}
	resp, err := json.Marshal(respMsg)
	if err != nil {
		ServerFail(ctx)
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.Write(resp)
}

func RespSucc(ctx *fasthttp.RequestCtx) {
	respMsg := map[string]int{
		"code": 0,
	}
	resp, _ := json.Marshal(respMsg)
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.Write(resp)
}
