package main

import (
	"fmt"
	"net"
)

func main() {
	iface, _ := net.InterfaceByName("eno1")

	addrs, _ := iface.Addrs()

	var ip string

	for _, addr := range addrs {
		ip, netIPNet, _ := net.ParseCIDR(addr.String())
		fmt.Println(ip)
		ip.To4()
		netIP := netIPNet.IP.To4()
		if netIP == nil {
			continue
		}
		// ip = netIP.String()

		break
	}
	fmt.Println(ip)
}
