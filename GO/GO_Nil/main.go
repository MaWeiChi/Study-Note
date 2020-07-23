package main

import "fmt"

type Foo struct {
	id   int
	name string
}

func assignValue() Foo {
	return Foo{id: 1, name: "abc"}
}

func main() {
	var foo *Foo
	fmt.Println(foo)
	fmt.Println(&foo)
	fmt.Println(assignValue())
	foo = new(Foo)
	if foo == nil {
		print("nil")
	}
	fmt.Println(*foo)
	*foo = assignValue()
	fmt.Println(foo)

}
