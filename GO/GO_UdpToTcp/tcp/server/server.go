package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	PORT := ":4001"
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	fmt.Printf("start server with %v\n", l.Addr())
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("connect with %v\n", c.RemoteAddr())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			if c, err = l.Accept(); err != nil {
				fmt.Println(err)
				return
			}
			continue
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
