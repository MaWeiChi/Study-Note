package main

import "fmt"

// func main() {
// 	go func() {
// 		fmt.Println("GO GO GO")
// 	}()
// 	time.Sleep(1 * time.Second)
// }

func basic() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO")
		c <- true
	}()
	<-c
}
