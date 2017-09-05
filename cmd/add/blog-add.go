package main

import (
	"blog/util"
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	name := flag.String("name", "", "The tag name")

	flag.Parse()

	input := bufio.NewScanner(os.Stdout)

	if len(*name) == 0 {
		fmt.Fprintln(os.Stdout, "Please enter the name of the tag:")
		input.Scan()
		*name = input.Text()
	}

	r := struct {
		Name string
	}{Name: *name}
	jsonVar, err := json.Marshal(r)
	util.CheckError(err)

	rep, err := http.Post("http://localhost:8080/tags", "application/json", bytes.NewBuffer(jsonVar))
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
