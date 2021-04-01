package main

import (
	"fmt"
	"regexp"
)

func main() {
	// re := regexp.MustCompile(`[^ $]`)

	// if !re.MatchString("") {
	// 	fmt.Println("not allow empty")
	// }
	// if !re.MatchString("    ") {
	// 	fmt.Println("not allow space")
	// }
	// if !re.MatchString("  dfas  ") {
	// 	fmt.Println("not allow with space1")
	// }
	// if !re.MatchString("dfas  ") {
	// 	fmt.Println("not allow with space2")
	// }
	// if !re.MatchString("   dfa") {
	// 	fmt.Println("not allow with space3")
	// }
	// if !re.MatchString(" ") {
	// 	fmt.Println("not allow one space")
	// }
	// if []byte{0}[0] == 0 {
	// 	fmt.Println(1)
	// }

	re := regexp.MustCompile(`\sversion\s`)
	s := "UC-8220-T-LX-EU-S (111) version V1.3"
	split := re.Split(s, 1)
	fmt.Println(split)
	for i := range split {
		fmt.Println(split[i])
	}
}
