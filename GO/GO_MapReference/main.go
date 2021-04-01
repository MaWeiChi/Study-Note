package main

import "fmt"

type Foo struct {
	name string
	id   int
}

var i int

func main() {
	m := make(map[string]interface{})
	m["f"] = w("123")
	fmt.Println(m["f"])
	m["f"] = w("456")
	fmt.Println(m["f"])
	w("789")
	fmt.Println(m["f"])

}

func w(t string) []Foo {
	i = i + 1
	return []Foo{name: t, id: i}
}
