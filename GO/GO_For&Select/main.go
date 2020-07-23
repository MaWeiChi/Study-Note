package main

import (
	"fmt"
	"time"
)

func main() {

	test()
}

func test() {
	i := 0

	for {
		fmt.Println("go")
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			fmt.Println(i)
			if i == 3 {
				continue
			}
			fmt.Println("inside the select: ")
		}
		fmt.Println("inside the for: ")
	}
}
