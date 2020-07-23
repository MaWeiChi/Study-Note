package main

import "fmt"

type Config struct {
	c
	cappend
}

type c struct {
	data Data
}

type Data struct {
	id   int
	name string
}
type cappend struct {
	data DataAppend
}
type DataAppend struct {
	sid       int
	lowername string
}

func main() {
	var first Config
	first.c.data.id = 1
	first.c.data.name = "fisrt"
	second := first
	second.c.data.id = 2
	fmt.Println(first)
	fmt.Println(second)
	first = second
	fmt.Println(first)

}
