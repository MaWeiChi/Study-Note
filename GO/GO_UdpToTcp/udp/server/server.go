package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	PORT := ":4001"
	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Println(err)
		fmt.Printf("start server with %v\n", addr)
		fmt.Print("-> ", string(buffer[0:n]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		_, err = connection.WriteToUDP([]byte(myTime), addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
