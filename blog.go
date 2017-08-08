package main

import (
	"blog/util"
	"github.com/julienschmidt/httprouter"
	"blog/api/v1.0/articles"
	"blog/api/v1.0/tags"
	"log"
	"net/http"
)

func main() {
	util.DbInit()

	defer util.Db.Close()

	err := util.Db.Ping()
	util.CheckError(err)

	router := httprouter.New()
	router.GET("/articles", articles.GetArticles)
	router.GET("/articles/:id", articles.GetArticle)
	router.GET("/tags/:id/articles", articles.GetArticleByTag)
	router.POST("/articles", articles.PostArticle)
	router.DELETE("/articles/:id", articles.DeleteArticle)
	router.PATCH("/articles/:id", articles.UpdateArticle)

	router.GET("/tags", tags.GetTags)
	router.GET("/tags/:id", tags.GetTag)
	router.GET("/articles/:id/tags", tags.GetTagByArticle)
	router.POST("/tags", tags.PostTag)
	router.DELETE("/tags/:id", tags.DeleteTag)

	log.Fatal(http.ListenAndServe(":8080", router))
}
