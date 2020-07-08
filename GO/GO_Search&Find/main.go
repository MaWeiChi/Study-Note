package main

import (
	"fmt"
	"reflect"
)

type abclist struct {
	list []abc123
}

type abc123 struct {
	word   string
	number int
}

func main() {
	var ls abclist
	ls.list = []abc123{{"A", 1}, {"B", 2}, {"c", 3}}
	// xs := []int{2, 4, 6, 8}
	// ys := []string{"C", "B", "K", "A"}

	// fmt.Println(Contains(xs, 8))
	fmt.Println(Contains(ls.list, 8))

	// fmt.Println(Find(ys, "C"))
}

func Contains(capsArray interface{}, item interface{}) bool {
	arr := reflect.ValueOf(capsArray)

	fmt.Print("arr: ")
	fmt.Println(arr)
	fmt.Print("arr.Kind(): ")
	fmt.Println(arr.Kind())

	if arr.Kind() != reflect.Slice && arr.Kind() != reflect.Array {
		return false
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {

			fmt.Print("arr.Index(i): ")
			fmt.Println(arr.Index(i))
			fmt.Print("arr.Index(i).Interface(): ")
			fmt.Println(arr.Index(i).Interface())

			return true
		}
	}
	return false
}

func Find(capsArray interface{}, item interface{}) int {
	arr := reflect.ValueOf(capsArray)
	if arr.Kind() != reflect.Slice && arr.Kind() != reflect.Array {
		return -1
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return i
		}
	}
	return -1
}
