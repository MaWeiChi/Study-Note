package main

import "fmt"

func main() {
	for l := 0; l < 5; l++ {
		defer fmt.Println("foo")
		fmt.Println(l)
	}
	var a int
	fmt.Println(a)
	if true {
		a := 2
		fmt.Println(a)
	}
	fmt.Println(a)
}
