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
func Handle(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		fmt.Println(string(ctx.Method()) + " - " + ctx.URI().String())

		next(ctx)
	})
}

func main() {
	router := fasthttprouter.New()
	router.GET("/ping", PingPong)

	// user
	router.GET("/api/users", Handle(controllers.GetAllUsers))
	router.POST("/api/login", Handle(controllers.Login))
	router.POST("/api/register", Handle(controllers.Register))

	// article
	router.GET("/api/articles", Handle(controllers.GetArticles))
	router.GET("/api/articlesTitles", Handle(controllers.GetArticlesTitles))
	router.GET("/api/article/:id", Handle(controllers.GetArticle))
	router.POST("/api/article", Handle(controllers.PostArticle))
	router.GET("/api/articles/tags", Handle(controllers.GetAllTags))

	//

	// xx

	fmt.Println("server is running at 9215...\n")
	log.Fatal(fasthttp.ListenAndServe(":9215", router.Handler))
}
