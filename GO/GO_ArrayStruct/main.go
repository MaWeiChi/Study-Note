package main

import "fmt"

type WifiUpdate struct {
	Type  string `json:"type"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Wan   bool   `json:"wan"`
	State string `json:"state"`
}

func main() {
	wifiUpdates := []WifiUpdate{{Type: "ethernets", Id: 1, Name: "wlan0", Wan: true, State: "connected"}, {Type: "ethernets", Id: 1, Name: "wlan0", Wan: true, State: "connected"}}
	fmt.Println(wifiUpdates)
}
