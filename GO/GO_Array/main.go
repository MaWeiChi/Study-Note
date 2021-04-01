package main

import (
	"fmt"
)

type Array struct {
	Array []string
}

func main() {
	// var arr Array
	// fmt.Println(len(arr.Array))
	// fmt.Println((arr.Array))

	xs := []int{2, 4, 6, 8}
	ys := []string{"C", "B", "K", "A"}
	fmt.Println(
		SliceIndex(len(xs), func(i int) bool { return xs[i] == 5 }),
		SliceIndex(len(xs), func(i int) bool { return xs[i] == 6 }),
		SliceIndex(len(ys), func(i int) bool { return ys[i] == "Z" }),
		SliceIndex(len(ys), func(i int) bool { return ys[i] == "A" }))
	xs2 := []int{2, 4, 6, 8}
	xs = append(xs, xs2...)

	fmt.Println(xs)

	fmt.Println(xs[1:])
	fmt.Println(xs[2:])
	fmt.Println(xs[3:])

}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
