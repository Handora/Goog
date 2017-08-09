package main

import (
	"blog/util"
	"net/http"
	"encoding/json"
	"fmt"
	"os"
	"flag"
)

func main() {
	page := flag.Int("page", 1, "The page of all articles you want to Get, default is 1")

	SomeArticles := struct {
		Code int
		Text string
		Body []util.Article
	}{}
	// get the request from api to get all tags
	req, err := http.Get(fmt.Sprintf("http://localhost:8080/articles?page=%d", *page))
	util.CheckError(err)

	defer req.Body.Close()

	// decode the response json and display it to the user
	json.NewDecoder(req.Body).Decode(&SomeArticles)
	fmt.Println(SomeArticles)
	fmt.Fprintln(os.Stdout, "All of you articles in page:")

	for _, i := range SomeArticles.Body {
		fmt.Fprintf(os.Stdout, "    %-5d%s\nTime: %v\n", i.Id, i.Title, i.Time)
		fmt.Fprintln(os.Stdout, i.Content)
	}
}
