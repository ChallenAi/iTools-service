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
	router.GET("/", Handle(controllers.Index))
	router.GET("/index.html", Handle(controllers.Index))

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

	// gist
	router.GET("/api/tips", Handle(controllers.GetGists))
	router.POST("/api/tip", Handle(controllers.AddGist))
	router.PUT("/api/tip/:id", Handle(controllers.EditGist))
	router.DELETE("/api/tip/:id", Handle(controllers.DeleteGist))

	// share
	router.GET("/api/shares", Handle(controllers.GetShares))
	router.POST("/api/share", Handle(controllers.AddShare))
	router.PUT("/api/share/:id", Handle(controllers.EditShare))
	router.DELETE("/api/share/:id", Handle(controllers.DeleteShare))

	// tag
	router.GET("/api/tags", Handle(controllers.GetTags))
	router.GET("/api/tagstree", Handle(controllers.GetTagsTree))
	router.POST("/api/tag", Handle(controllers.AddTag))
	router.PUT("/api/tag", Handle(controllers.EditTag))
	router.DELETE("/api/tag", Handle(controllers.DeleteTag))

	// bookmark
	router.GET("/api/bookmarks", Handle(controllers.GetBookmarks))
	router.DELETE("/api/bookmark/:id", Handle(controllers.DeleteBookmark))

	// pass
	router.GET("/api/passwords", Handle(controllers.GetPasswords))
	router.POST("/api/pass", Handle(controllers.AddPass))
	router.PUT("/api/pass/:id", Handle(controllers.EditPass))
	router.DELETE("/api/pass/:id", Handle(controllers.DeletePass))

	router.NotFound = controllers.NotFound

	fmt.Println("server is running at " + conf.Port + "...")
	log.Fatal(fasthttp.ListenAndServe(":"+conf.Port, router.Handler))
}
