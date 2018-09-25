package controllers

import (
	"bytes"
	// "fmt"
	"github.com/ChallenAi/iTools-service/utils"
	"github.com/valyala/fasthttp"
	"io/ioutil"
)

func Index(ctx *fasthttp.RequestCtx) {
	const htmlPath = "./static/index.html"
	body, _ := ioutil.ReadFile(htmlPath)
	ctx.Write(body)
	ctx.Response.Header.SetContentType("text/html")
	// fasthttp.FSHandler("./static/", 1)(ctx)
}

func Pong(ctx *fasthttp.RequestCtx) {
	utils.RespData(ctx, "pong")
}

func NotFound(ctx *fasthttp.RequestCtx) {
	path := ctx.Path()
	if bytes.HasPrefix(path, []byte("/static/css/main.d19c0ba4.css")) {
		const cssPath = "./static/css/main.d19c0ba4.css"
		body, _ := ioutil.ReadFile(cssPath)
		ctx.Write(body)
		ctx.Response.Header.SetContentType("text/css")
	} else if bytes.HasPrefix(path, []byte("/static/js/main.d4bd205e.js")) {
		const jsPath = "./static/js/main.d4bd205e.js"
		body, _ := ioutil.ReadFile(jsPath)
		ctx.Write(body)
		ctx.Response.Header.SetContentType("text/javascript")
	} else if bytes.HasPrefix(path, []byte("/service-worker.js")) {
		const workerPath = "./static/service-worker.js"
		body, _ := ioutil.ReadFile(workerPath)
		ctx.Write(body)
		ctx.Response.Header.SetContentType("text/javascript")
	} else {
		ctx.SetStatusCode(404)
	}
}
