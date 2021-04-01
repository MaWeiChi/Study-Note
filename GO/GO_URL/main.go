package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println((&url.URL{Path: "1:7.4p1-10+deb9u7"}).String())
	fmt.Println(url.QueryEscape("1:7.4p1-10+deb9u7"))
}
