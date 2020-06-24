package main

import "fmt"

type Config struct {
	id   int
	name string
}

func main() {
	var first Config
	first.id = 1
	first.name = "fisrt"
	second := first
	second.id = 2
	fmt.Println(first)
	fmt.Println(second)
	first = second
	fmt.Println(first)

}
