package main

import "fmt"

func main() {
	var a int
	fmt.Println(a)
	if true {
		a := 2
		fmt.Println(a)
	}
	fmt.Println(a)
}
