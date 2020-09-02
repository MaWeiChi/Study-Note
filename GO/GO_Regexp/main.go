package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`[^ $]`)

	if !re.MatchString("") {
		fmt.Println("not allow empty")
	}
	if !re.MatchString("    ") {
		fmt.Println("not allow space")
	}
}
