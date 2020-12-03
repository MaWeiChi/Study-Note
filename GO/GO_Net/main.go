package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Interfaces()
	if err != nil {
		panic(err)

	}
	for _, f := range l {
		if f.Flags&net.FlagUp > 0 {
			fmt.Printf("%s is up\n", f.Name)
		}
	}
}
