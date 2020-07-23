package main

import "fmt"

type Foo struct {
	id   int
	name string
}

func main() {
	foo := []Foo{{id: 1, name: "aaa"}, {id: 2, name: "bbb"}, {id: 3, name: "ccc"}}
	fmt.Println(foo)
	for _, f := range foo {
		f.id = f.id + 1
	}
	fmt.Println(foo)
	for i, _ := range foo {
		foo[i].id = foo[i].id + 1
	}
	fmt.Println(foo)
}
