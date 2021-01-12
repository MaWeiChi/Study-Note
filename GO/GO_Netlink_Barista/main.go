package main

import (
	"fmt"
	"net"
	"time"

	"barista.run/base/watchers/netlink"
)

func main() {
	list, _ := net.Interfaces()
	for _, l := range list {
		l.Name
	}
	scEno1 := netlink.ByName("eth0")
	eno1 := scEno1.Get()
	fmt.Println(eno1.Name)
	fmt.Println(eno1.HardwareAddr)
	fmt.Println(eno1.State)
	fmt.Println(eno1.IPs)
	for {
		select {
		case <-scEno1.C:
			eno1 = scEno1.Get()
			fmt.Println(eno1.State)

		case <-time.After(time.Duration(2 * time.Second)):
			fmt.Print("detecting: ")
			fmt.Print("current ")
			fmt.Println(eno1.Name)
			fmt.Println(eno1.IPs)

			fmt.Println(eno1.State)
		}

	}
}
