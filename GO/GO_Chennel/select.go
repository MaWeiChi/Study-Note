package main

import (
	"fmt"
)

var Dws string = "qwd"

func selectdemo() {
	ch := make(chan int, 3)
	ch <- 1
	select {
	case ch <- 2:
		ch <- 3
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)

	case ch <- 3:
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)

	default:
		fmt.Println("channel blocking")
	}
}
