package main

import (
	"fmt"
	"github.com/ChallenAi/iTools-service/conf"
	"github.com/ChallenAi/iTools-service/controllers"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

// middleware
func Handle(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		fmt.Println(string(ctx.Method()) + " - " + ctx.URI().String())

		next(ctx)
	})
}

func main() {
	db := models.InitDB()
	defer db.Close()

	router := fasthttprouter.New()
	router.GET("/ping", Handle(controllers.Pong))

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

	router.GET("/api/tips", Handle(controllers.GetGists))

	router.NotFound = controllers.NotFound

	router.ServeFiles("/img/*filepath", "static/img")

	fmt.Println("server is running at " + conf.Port + "...")
	log.Fatal(fasthttp.ListenAndServe(":"+conf.Port, router.Handler))
}
