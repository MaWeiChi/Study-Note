package main

import "fmt"

func main() {
	a := []string{"Hello1", "Hello2", "Hello3"}
	fmt.Println(a)
	// [Hello1 Hello2 Hello3]
	a = append(a[:1], a[2:]...)
	fmt.Println(a)
	// [Hello2 Hello3]
}
