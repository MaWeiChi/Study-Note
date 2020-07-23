package main

import "fmt"

func main() {
	var s string
	s = "complete"
	ConnectedConfirm(&s)
	fmt.Println(s)
}

func ConnectedConfirm(status *string) {
	if *status == "complete" {
		*status = "connected"
	}
}
