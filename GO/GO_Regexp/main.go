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
	if !re.MatchString("  dfas  ") {
		fmt.Println("not allow with space1")
	}
	if !re.MatchString("dfas  ") {
		fmt.Println("not allow with space2")
	}
	if !re.MatchString("   dfa") {
		fmt.Println("not allow with space3")
	}
	if !re.MatchString(" ") {
		fmt.Println("not allow one space")
	}
	if []byte{0}[0] == 0 {
		fmt.Println(1)
	}
}
