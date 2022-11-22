package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type ServerItem struct {
	Protocol string `json:"protocol, validate:"oneof=tcp udp"`
	Port     int    `json:"port, validate:"min=0,max=65536"`
}

type ClientItem struct {
	Protocol string `json:"protocol, validate:"oneof=tcp udp"`
	Port     int    `json:"port, validate:"min=0,max=65536"`
	Host     string `json:"host, validate:"ipv4"`
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	createTCPServer("8443")
	createUDPServer("4001")
	createTCPClient("127.0.0.1:8443")
	createUDPClient("127.0.0.1:1234")
}

func createConnect(serverOpt ServerItem, clientOpt ClientItem, ctx context.Context) {
	var serverConn, clientConn net.Conn
	var err error

	defer serverConn.Close()
	defer clientConn.Close()

	if serverOpt.Protocol == "tcp" {
		s, err := net.ResolveTCPAddr("tcp", strconv.Itoa(serverOpt.Port))
		if err != nil {
			fmt.Println(err)
			return
		}
		l, err := net.ListenTCP("tcp", s)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer l.Close()
		serverConn, err = l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if serverOpt.Protocol == "udp" {
		s, err := net.ResolveUDPAddr("udp4", strconv.Itoa(serverOpt.Port))
		if err != nil {
			fmt.Println(err)
			return
		}

		serverConn, err = net.ListenUDP("udp4", s)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	addr := clientOpt.Host + ":" + strconv.Itoa(clientOpt.Port)
	if clientOpt.Protocol == "tcp" {
		clientConn, err = net.Dial("tcp", addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		clientConn.Close()
	} else if clientOpt.Protocol == "udp" {
		s, err := net.ResolveUDPAddr("udp4", addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		clientConn, err = net.DialUDP("udp4", nil, s)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	go func() {
		for {
			// server -> client
			io.Copy(clientConn, serverConn)
			if err == io.EOF {
				fmt.Println(err)
				serverConn, err = l.Accept()

			}

		}
	}()
	go func() {
		io.Copy(serverConn, clientConn)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
}

func createTCPServer(port string) (err error) {

	PORT := ":" + port
	s, err := net.ResolveTCPAddr("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	buffer := make([]byte, 1024)

	for {
		c.Read(buffer)
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return err
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return err
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}

func createUDPServer(port string) (err error) {
	PORT := ":" + port

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
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Print("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return err
		}

		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
}

func createTCPClient(addr string) (err error) {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}

func createUDPClient(addr string) (err error) {
	// go run udpC.go 127.0.0.1:1234
	// addr example: 127.0.0.1:1234
	s, err := net.ResolveUDPAddr("udp4", addr)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err = c.Write(data)
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
}
