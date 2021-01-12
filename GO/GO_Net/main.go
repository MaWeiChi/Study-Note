package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

// func main() {
// 	l, err := net.Interfaces()
// 	if err != nil {
// 		panic(err)

// 	}
// 	for _, f := range l {
// 		fmt.Println("=========")
// 		fmt.Println(f.Name)
// 		fmt.Println(f.Flags.String())
// 		fmt.Println(f.HardwareAddr.String())
// 		fmt.Println(f.MTU)
// 		fmt.Println("=========")

// 		// if f.Flags&net.FlagUp > 0 {
// 		// 	fmt.Printf("%s is up", f.Name)
// 		// 	a, _ := f.Addrs()
// 		// 	fmt.Println(a)

// 		// }
// 	}

// w, err := net.InterfaceByName("eno1")
// if err != nil {
// 	panic(err)

// }
// fmt.Println(w.Name)

// if w.Flags&net.FlagUp > 0 {
// 	fmt.Printf("%s is up", w.Name)
// 	a, _ := w.Addrs()
// 	fmt.Println(a)

// }
// addrs, err := w.Addrs()
// fmt.Println(addrs)
// for range addrs {
// 	fmt.Println(1)

// }
// // a, err := net.InterfaceAddrs()
// // if err != nil {
// // 	panic(err)

// // }
// // for _, f := range a {
// // 	fmt.Println(f.String())
// // 	fmt.Println(f.Network())

// // }
// ifaces, _ := net.Interfaces()
// // handle err
// for _, i := range ifaces {
// 	if i.Name != "eno1" {
// 		continue
// 	}
// 	fmt.Println(i.Name)
// 	addrs, _ := i.Addrs()
// 	// handle err
// 	// for _, addr := range addrs {
// 	// 	var ip net.IP
// 	// 	switch v := addr.(type) {
// 	// 	case *net.IPNet:
// 	// 		fmt.Println("IPNet")
// 	// 		ip = v.IP
// 	// 		fmt.Println(ip)
// 	// 		addr := ip.To4()
// 	// 		fmt.Println(v.Mask)
// 	// 		sz, _ := net.IPv4Mask(addr[0], addr[1], addr[2], addr[3]).Size()
// 	// 		fmt.Println(sz)
// 	// 	case *net.IPAddr:
// 	// 		fmt.Println("IPAddr")
// 	// 		ip = v.IP
// 	// 	}
// 	for _, addr := range addrs {
// 		var ip net.IP
// 		// var mask net.IPMask
// 		switch v := addr.(type) {
// 		case *net.IPNet:
// 			fmt.Println("IPNet")

// 			ip = v.IP
// 			// mask = v.Mask
// 		case *net.IPAddr:
// 			ip = v.IP
// 			// mask = ip.DefaultMask()
// 		}
// 		if ip == nil {
// 			continue
// 		}
// 		addr1 := ip.To4()
// 		if addr1 == nil {
// 			continue
// 		}
// 		fmt.Println(addr1)

// 		fmt.Println(len(addr1))

// 		sz, _ :=net.IPv4Mask(addr1[0], addr1[1], addr1[2], addr1[3]).Size()
// 		// fmt.Println(mask)

// 		fmt.Println(sz)
// 		// process IP address
// 	}
// }

// iface, err := net.InterfaceByName("eno1")
// if err != nil {
// 	log.Println("Unable to find interface eth0, trying en0")
// 	iface, err = net.InterfaceByName("eno1")
// }
// fmt.Println(iface.HardwareAddr.String())
// // addrs, err := mgmtInterface.Addrs()
// // if err != nil {
// // 	log.Println("interface has no address")
// // }
// addrs, err := iface.Addrs()
// if err != nil {
// 	log.Println("interface has no address")
// }
// for _, addr := range addrs {
// 	fmt.Println(addr)
// 	iip, mmask, err := net.ParseCIDR(addr.String())
// 	ip := iip.To4()
// 	if ip == nil {
// 		continue
// 	}
// 	fmt.Println(iip.String())
// 	fmt.Println(mmask.Mask)
// 	mask := mmask.Mask
// 	cleanMask := fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3])
// 	fmt.Println(cleanMask)
// 	fmt.Println(err)
// 	// fmt.Println(GetSubnet(iip.String(), cleanMask))
// 	fmt.Println(GetBroadcast(iip.String(), cleanMask))

// }
// for _, addr := range addrs {
// 	fmt.Println(addr)
// 	iip, mmask, err := net.ParseCIDR(addr.String())
// 	fmt.Println(iip)
// 	fmt.Println(mmask.Mask)
// 	fmt.Println(err)
// 	var ip net.IP
// 	var mask net.IPMask

// 	switch v := addr.(type) {

// 	case *net.IPNet:
// 		ip = v.IP
// 		mask = v.Mask
// 		fmt.Println(ip)

// 	case *net.IPAddr:
// 		ip = v.IP
// 		mask = ip.DefaultMask()
// 	}
// 	if ip == nil {
// 		continue
// 	}
// 	ip = ip.To4()
// 	if ip == nil {
// 		continue
// 	}
// 	fmt.Println(mask)
// 	// create the netmask
// 	cleanMask := fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3])
// 	// addr := ip.To4()
// 	// sz, _ := net.IPV4Mask(addr[0], addr[1], addr[2], addr[3]).Size()
// 	fmt.Println(cleanMask)
// 	fmt.Println(GetMask("192.168.3.97/24"))
// }
// }

// ipnet := net.IPNet{
//     IP:   net.ParseIP(ip.String()),
//     Mask: net.IPv4Mask(m[0], m[1], m[2], m[3]),
// }

// iface, err := net.InterfaceByName(name)

func GetMask(netmask string) string {
	mask := addrStringToBytes(netmask)
	if len(mask) != 4 {
		panic("ipv4Mask: len must be 4 bytes")
	}
	return fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3])
}

func addrStringToBytes(ip string) []byte {
	addrs := make([]byte, 4)
	ipstrs := strings.Split(ip, ".")
	for i := range ipstrs {
		num, _ := strconv.Atoi(ipstrs[i])
		addrs[i] = byte(num)
	}
	return addrs
}

func GetSubnet(ip string, mask string) string {
	ipaddrs := addrStringToBytes(ip)
	maskaddrs := addrStringToBytes(mask)
	networkQuards := []string{}
	networkQuards = append(networkQuards, fmt.Sprintf("%d", ipaddrs[0]&maskaddrs[0]))
	networkQuards = append(networkQuards, fmt.Sprintf("%d", ipaddrs[1]&maskaddrs[1]))
	networkQuards = append(networkQuards, fmt.Sprintf("%d", ipaddrs[2]&maskaddrs[2]))
	networkQuards = append(networkQuards, fmt.Sprintf("%d", ipaddrs[3]&maskaddrs[3]))
	return strings.Join(networkQuards, ".")
}

func GetBroadcast(ip string, netmask string) string {
	addrs := addrStringToBytes(ip)
	mask := addrStringToBytes(netmask)
	broadcast := net.IP(make([]byte, 4))
	for i := range addrs {
		broadcast[i] = addrs[i] | ^mask[i]
	}
	return fmt.Sprintf("%d.%d.%d.%d", broadcast[0], broadcast[1], broadcast[2], broadcast[3])
}

type AdapterWithIPNets struct {
	Name   string
	IPNets []net.IPNet
}

func Adapters() ([]AdapterWithIPNets, error) {
	var awins []AdapterWithIPNets
	ai, err := getAdapterList()
	if err != nil {
		return nil, err
	}
	for ; ai != nil; ai = ai.Next {
		name := bytePtrToString(&ai.AdapterName[0])
		awin := AdapterWithIPNets{Name: name}
		iai := &ai.IpAddressList
		for ; iai != nil; iai = iai.Next {
			ip := net.ParseIP(bytePtrToString(&iai.IpAddress.String[0]))
			mask := parseIPv4Mask(bytePtrToString(&iai.IpMask.String[0]))
			awin.IPNets = append(awin.IPNets, net.IPNet{IP: ip, Mask: mask})
		}
		awins = append(awins, awin)
	}
	return awins, nil
}

func parseIPv4Mask(ipStr string) net.IPMask {
	ip := net.ParseIP(ipStr).To4()
	return net.IPv4Mask(ip[0], ip[1], ip[2], ip[3])
}

// https://github.com/golang/go/blob/go1.4.1/src/net/interface_windows.go#L13-L20
func bytePtrToString(p *uint8) string {
	a := (*[10000]uint8)(unsafe.Pointer(p))
	i := 0
	for a[i] != 0 {
		i++
	}
	return string(a[:i])
}

// copied from https://github.com/golang/go/blob/go1.4.1/src/net/interface_windows.go#L22-L39
func getAdapterList() (*ai.IpAdapterInfo, error) {
	b := make([]byte, 1000)
	l := uint32(len(b))
	a := (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
	// TODO(mikio): GetAdaptersInfo returns IP_ADAPTER_INFO that
	// contains IPv4 address list only. We should use another API
	// for fetching IPv6 stuff from the kernel.
	err := syscall.GetAdaptersInfo(a, &l)
	if err == syscall.ERROR_BUFFER_OVERFLOW {
		b = make([]byte, l)
		a = (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
		err = syscall.GetAdaptersInfo(a, &l)
	}
	if err != nil {
		return nil, os.NewSyscallError("GetAdaptersInfo", err)
	}
	return a, nil
}

func main() {
	awins, err := Adapters()
	if err != nil {
		panic(err)
	}
	for _, awin := range awins {
		fmt.Printf("name=%s\n", awin.Name)
		for _, ipnet := range awin.IPNets {
			fmt.Printf("addr=%s, mask=%s\n", ipnet.IP, ipnet.Mask)
		}
		fmt.Println()
	}
}
