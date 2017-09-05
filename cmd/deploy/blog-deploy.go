package main

import (
	"blog/cmd/cmdUtility"
	"blog/util"
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	// []int Interface for us to serve the flag for several arguments
	var tags cmdUtility.ArrayFlags

	var err error

	// go flag, useful for cli programming
	article := flag.String("article", "", "The article you want to upload(should end with '.md')")
	title := flag.String("title", "", "The Article title you want to list")
	flag.Var(&tags, "tag", "The tags of your article, may be multiple arguments as -tag 1 -tag 2")

	flag.Parse()

	input := bufio.NewScanner(os.Stdin)

	// if article flag is empty, reenter the article
	if len(*article) == 0 {
		cmdUtility.GetArticleContent(article, input)
	}

	// if the file is not ends with md, we should panic
	if !strings.HasSuffix(*article, ".md") {
		fmt.Fprintln(os.Stdout, "the Article flag must ends with '.md'")
		os.Exit(1)
	}

	if len(*title) == 0 {
		cmdUtility.GetArticleTitle(title, input)
	}

	// if the tag is empty, we should recommend user to rewrite it
	if len(tags) == 0 {
		cmdUtility.GetArticleTags(tags, input)
	}

	// read the file
	md, err := ioutil.ReadFile(*article)
	util.CheckError(err)

	// construct the intro
	re, err := regexp.Compile("<!--intro-->")
	util.CheckError(err)

	index := re.FindStringIndex(string(md))

	var intro string = ""
	if index != nil {
		intro = string(md)[:index[0]]
	}

	// construct the post body for posting articles
	r := struct {
		Title   string
		Intro   string
		Content string
		Tag     []int
	}{Title: *title, Intro: intro, Content: string(md), Tag: tags}
	jsonVar, err := json.Marshal(r)
	util.CheckError(err)

	// get the response
	rep, err := http.Post("http://localhost:8080/articles", "application/json", bytes.NewBuffer(jsonVar))
	util.CheckError(err)

	// test the response's status code
	if rep.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stdout, "response return code %d\n", rep.StatusCode)
		os.Exit(1)
	}

	// if all success, just say success and be happy!
	fmt.Fprintln(os.Stdout, "success!")
	os.Exit(0)
}
