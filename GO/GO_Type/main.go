package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i32 int32
	i32 = 1
	println(i32)
	println(string(i32))
	var i int
	i = 1
	println(strconv.Itoa(i))
	fmt.Printf("network Disconnect Code: %d\n", int((i32)))
	fmt.Printf("network Disconnect Code: %s\n", strconv.Itoa(int(i32)))
}
