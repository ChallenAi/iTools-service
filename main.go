package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

type OkResp struct {
	Code    int
	Message string
}

func Index(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.WriteString("hello")
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)

	log.Fatal(fasthttp.ListenAndServe(":9215", router.Handler))
}
