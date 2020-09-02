package main

import "fmt"

func main() {
	if true {
		return
	}
	if true {
		defer fmt.Println("1")
		defer fmt.Println("0")
	}

	fmt.Println(2)
}
