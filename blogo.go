package main

import (
	"blog/api/v1.0/articles"
	"blog/api/v1.0/tags"
	"blog/util"
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	// get the config file from input
	config := flag.String("config", "", "must be a config json file")
	flag.Parse()

	// if no input, we just print the usage and quit
	if len(*config) == 0 {
		fmt.Println("Usage: blog -config={config.json}")
		os.Exit(1)
	}

	// init the util.Db for later use
	util.DbInit(*config)

	// dynamic close the Db using defer
	defer util.Db.Close()

	// test whether connect the database
	err := util.Db.Ping()
	util.CheckError(err)

	// creating the restful web server using github.com/julienschmidt/httprouter
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

	// Now start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
