package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// global variables
var destUrl string

/**
download images
*/
func downloadImages(i int, e *colly.HTMLElement) {
	classname := e.Attr("class")
	if classname != "image" {
		return
	}
	srcRef := e.Attr("href")
	// fmt.Println(srcRef)
	segments := strings.Split(srcRef, "/")
	fileName := segments[len(segments)-1]
	fullurl := destUrl + "#/media/" + fileName
	fmt.Println(fullurl)
	res, _ := http.Get(fullurl)
	savedPath := "Study-Note/GO/GO_Colly&Image/" + fileName
	// TODO create dirs before create file
	f, err := os.Create(savedPath)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
	defer res.Body.Close()
	defer f.Close()

	// fmt.Println(destUrl + srcRef)
}

func main() {
	c := colly.NewCollector()

	c.OnHTML("div[id]", func(e *colly.HTMLElement) {

		className := e.Attr("id")
		// fmt.Println(className)
		// e.ForEach("img[src]", downloadImages)
		if className == "mw-content-text" {
			e.ForEach("a[href]", downloadImages)

		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	destUrl = "https://zh.wikipedia.org/wiki/Google"
	c.Visit(destUrl)
}
