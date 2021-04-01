package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"regexp"
	"strings"
)

type EthernetInfo struct {
	DisplayName string   `json:"displayName"`
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Enable      bool     `json:"enable"`
	Status      string   `json:"status"`
	IP          string   `json:"ip"`
	Netmask     string   `json:"netmask"`
	Subnet      string   `json:"subnet"`
	Broadcast   string   `json:"broadcast"`
	Gateway     string   `json:"gateway"`
	DNS         []string `json:"dns"`
	MAC         string   `json:"mac"`
	Speed       string   `json:"speed"`
	// // mock wan&dhcp
	// EnableDHCP bool `json:"enableDhcp"`
	// Wan        bool `json:"wan"`
}

func main() {
	var ethernets []EthernetInfo
	var ethMapArray map[string]int
	ethMapArray = make(map[string]int)
	reEn := regexp.MustCompile(`eno[0-9]`)
	reEw := regexp.MustCompile(`wlp[0-9]s0`)
	net, _ := net.Interfaces()
	for _, n := range net {
		if reEn.MatchString(n.Name) || reEw.MatchString(n.Name) {
			ethernets = append(ethernets, EthernetInfo{Name: n.Name, DisplayName: n.Name})
			ethMapArray[n.Name] = len(ethernets) - 1
		}
		fmt.Println(n.Name)
	}
	// out, _ := exec.Command("cat", "/etc/network/interfaces").Output()
	// for _, l := range strings.Split(string(out), "\n") {
	// 	line := strings.TrimSpace(l)
	// 	splits := strings.Split(line, " ")
	// 	if splits[0] != "iface" {
	// 		continue
	// 	}
	// 	if reEn.MatchString(splits[1]) || reEw.MatchString(splits[1]) {
	// 		ethernets = append(ethernets, EthernetInfo{Name: splits[1], DisplayName: splits[1]})
	// 		ethMapArray[splits[1]] = len(ethernets) - 1
	// 	}
	// }
	fmt.Println(ethernets)
	fmt.Println(len(ethernets))
	for i, eth := range ethernets {
		path := fmt.Sprintf("/sys/class/net/%s/speed", eth.Name)
		fmt.Print(path)
		// out, _ := exec.Command("cat", path).Output()
		out, err := ioutil.ReadFile(path)
		if err != nil {

			continue
		}
		fmt.Println(out)

		ethernets[i].Speed = fmt.Sprintf("%sMb/s", strings.TrimSpace(string(out)))
		fmt.Println(ethernets[i].Name)
		fmt.Println(ethernets[i].Speed)
		fmt.Println(ethernets[i].Speed)

	}
	fmt.Println(ethernets)
}
