package main

import (
	"fmt"
)

type ApServerList struct {
	State   string                   `json:"state"`
	List    map[string]ApServerEntry `json:"list"`
	IP      string                   `json:"ip"`
	Netmask string                   `json:"netmask"`
	Count   int                      `json:"count"`
}

type ApServerEntry struct {
	ID       int    `json:"id"`
	Ssid     string `json:"ssid"`
	Bssid    string `json:"bssid"`
	PSK      string `json:"psk"`
	Signal   int64  `json:"signal"`
	Band     string `json:"band"`
	KeyMgmt  string `json:"keyMgmt"`
	Current  bool   `json:"current"`
	Priority int    `json:"priority"`
}

func main() {
	var wifi ApServerList
	wifi.List = make(map[string]ApServerEntry)
	// wifi.List["s"] = ApServerEntry{}
	sside := "moxa"
	entry := ApServerEntry{1, "ssid", "bssid", "psk", -30, "5G", "wpa", true, 2}
	// if _, ok := wifi.List[sside]; !ok {
	// 	wifi.List[sside] = ApServerEntry{}
	// }
	wifi.List[sside] = entry
	// wifi.List[sside] = append(wifi.List[sside], entry)
	// wifi.List["b"] = "ball"
	fmt.Println(wifi.List)
	fmt.Println(len(wifi.List))
	fmt.Println(wifi)
	// fmt.Println(len(wifi.List["moxa"]))
	// fmt.Println(len(wifi.List["m"]))
	// fmt.Println(wifi.List["b"])
}
