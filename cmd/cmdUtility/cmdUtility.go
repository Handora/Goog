package cmdUtility

import (
	"blog/util"
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
	an implementation for flag to receive several arguments
*/
type ArrayFlags []int

// Value ...
func (i *ArrayFlags) String() string {
	return fmt.Sprint(*i)
}

// Set 方法是flag.Value接口, 设置flag Value的方法.
// 通过多个flag指定的值， 所以我们追加到最终的数组上.
func (i *ArrayFlags) Set(value string) error {
	v, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	*i = append(*i, v)
	return nil
}

// if the article flag is empty, we should recommend user to rewrite it
func GetArticleContent(article *string, input *bufio.Scanner) {
	// if the file flag is empty, we should recommend user to rewrite it
	fmt.Fprintln(os.Stdout, "Please enter the file path of which you want to upload(should end with '.md'):")
	input.Scan()
	*article = input.Text()
}

// if the title flag is empty, we should recommend user to rewrite it
func GetArticleTitle(title *string, input *bufio.Scanner) {
	// if the title flag is empty, we should recommend user to rewrite it
	fmt.Fprintln(os.Stdout, "Please enter the title of the articles:")
	input.Scan()
	*title = input.Text()
}

// if the tag flag is empty, we should recommend user to rewrite it
func GetArticleTags(tags ArrayFlags, input *bufio.Scanner) {
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
	fmt.Fprintln(os.Stdout, "Please choose whatever tag you like(just enter if nothing):")

	for _, i := range allTags.Body {
		fmt.Fprintf(os.Stdout, "    %-5d%s\n", i.Id, i.Name)
	}

	// get the input
	input.Scan()
	chose := input.Text()

	// and split it into []string
	w := strings.FieldsFunc(chose, func(r rune) bool {
		switch r {
		case ',', ';', ' ', '\t', '.', '\n':
			return true
		default:
			return false
		}
	})

	// convert it into []int and complete the tags
	for _, i := range w {
		n, err := strconv.Atoi(i)
		util.CheckError(err)
		tags = append(tags, n)
	}

}
