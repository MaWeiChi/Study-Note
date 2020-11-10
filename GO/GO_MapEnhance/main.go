package main

import "fmt"

func main() {
	fmt.Println(mapEthernet())
}

func mapEthernet() map[int]NetworkLayerEntry {
	var ethMap map[int]NetworkLayerEntry
	ethMap = make(map[int]NetworkLayerEntry)
	var N NetworkLayerEntry
	N.Ipsetting.EnableDhcp = true
	N.Ipsetting.Ip = "192.168.3.100"
	N.Ipsetting.Netmask = "255.255.252.0"
	N.Ipsetting.Gateway = "192.168.3.1"
	N.Ipsetting.Dns = []string{"8.8.8.8"}
	N.Ipsetting.EthernetId = 1

	ethMap[1] = N
	M := ethMap[1]
	M.Ipsetting.Netmask = "255.255.252.0"
	ethMap[1] = M
	return ethMap
}
