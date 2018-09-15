package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func PingPong(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.WriteString("pong")
}

func main() {
	router := fasthttprouter.New()
	router.GET("/ping", PingPong)

	fmt.Println("server is running at 9215...\n")
	log.Fatal(fasthttp.ListenAndServe(":9215", router.Handler))
}
