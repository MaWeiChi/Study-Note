package main

import "fmt"

func main() {
	var a uint
	var b uint
	a = 1
	b = 2
	if a < b {
		fmt.Println(b)
	}
	c := a - b
	fmt.Println(c)
	//18446744073709551615
	d := -1
	fmt.Println(uint(d))
	//18446744073709551615
}
