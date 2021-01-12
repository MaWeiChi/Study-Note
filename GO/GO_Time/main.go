package main

import (
	"fmt"
	"time"
)

var t string

func main() {
	fmt.Println(time.Now())
	t = time.Now().String()
	fmt.Println(t)
	fmt.Println(time.Now().Format(time.RFC3339))
}
