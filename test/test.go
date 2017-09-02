package main

import (
	"net/http"
	"fmt"
)

func main() {
	fmt.Printf("%d", 15/4)
	Test()
}

func Test() {
	_, err := http.Get("www.baidu.com")
	if err != nil {
		fmt.Printf("%v", err)
	}
}