package main

import (
	"fmt"

	s "github.com/teris-io/shortid"
)

func main() {
	fmt.Println(s.Generate())
}
