package main

import "fmt"

type Array struct {
	Array []string
}

func main() {
	var arr Array
	fmt.Println(len(arr.Array))
	fmt.Println((arr.Array))
}
