package main

import "fmt"

func main() {
	Test(nil)
}

func Test(data []byte) {
	if len(data) == 0 {
		fmt.Println("data is empty")
	} else {
		fmt.Println(string(data))
	}
}
