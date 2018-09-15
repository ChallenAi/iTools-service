package main

import (
	"fmt"
	"github.com/ChallenAi/iTools-service/controllers"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func PingPong(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.WriteString("pong")
}

// Maybe there's a better way to realize a middleware
func Logger(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		fmt.Println(string(ctx.Method()) + " - " + ctx.URI().String())

		next(ctx)
	})
}

func main() {
	router := fasthttprouter.New()
	router.GET("/ping", PingPong)

	router.GET("/api/users", Logger(controllers.GetAllUsers))
	router.POST("/api/login", Logger(controllers.Login))
	router.POST("/api/register", Logger(controllers.Register))

	fmt.Println("server is running at 9215...\n")
	log.Fatal(fasthttp.ListenAndServe(":9215", router.Handler))
}
