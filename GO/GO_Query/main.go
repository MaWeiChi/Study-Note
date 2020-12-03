package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	fmt.Println(doc.Contents())
	for _, item := range doc.Contents().Nodes {
		fmt.Println("{")
		fmt.Println(item.Type)
		fmt.Println(item.DataAtom)
		fmt.Println(item.Data)
		fmt.Println(item.Namespace)
		fmt.Println(item.Attr)
		fmt.Println("}")
		// fmt.Println(item.Attr)
		// fmt.Println(item.Data)
	}
	// doc.Find("article").Each(func(i int, s *goquery.Selection) {
	// 	// For each item found, get the band and title
	// 	band := s.Find("a").Text()
	// 	title := s.Find("i").Text()
	// 	fmt.Printf("Review %d: %s - %s\n", i, band, title)
	// })
}

func main() {
	fmt.Println(123)
	ExampleScrape()
}
