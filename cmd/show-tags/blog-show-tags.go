package main

import (
	"blog/util"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	allTags := struct {
		Code int
		Text string
		Body []util.Tag
	}{}
	// get the request from api to get all tags
	req, err := http.Get("http://localhost:8080/tags")
	util.CheckError(err)

	defer req.Body.Close()

	// decode the response json and display it to the user
	json.NewDecoder(req.Body).Decode(&allTags)
	fmt.Fprintln(os.Stdout, "All of you current tags:")

	for _, i := range allTags.Body {
		fmt.Fprintf(os.Stdout, "    %-5d%s\n", i.Id, i.Name)
	}
}
