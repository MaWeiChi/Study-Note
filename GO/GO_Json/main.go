package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type TimeEntry struct {
	Type string `json:"type"`
	TimeReadOnlyEntry
	Time string `json:"time" validate:"omitempty,min=20,max=35,timerfc3339"`
	TimeConfigEntry
}

type TimeConfigEntry struct {
	Timezone string   `json:"timezone" validate:"omitempty,min=1,max=64,alpha|contains=/"`
	Ntp      NtpEntry `json:"ntp"`
}

type TimeReadOnlyEntry struct {
	LastUpdateTime string `json:"lastUpdateTime" validate:"omitempty"`
}

type NtpEntry struct {
	Enable      bool   `json:"enable"`
	Source      string `json:"source" validate:"omitempty,oneof=gps timeserver"`
	Server      string `json:"server" validate:"omitempty,min=2,max=253,ipv4|fqdn"`
	GPSLongJump bool   `json:"gpsLongJump"`
	Interval    int    `json:"interval" validate:"omitempty,min=60,max=2592000"`
}

type AddEntry struct {
	Ssid     string `json:"ssid"`
	Password string `json:"passwrod"`
	Priority int    `json:"priority"`
}

func main() {

	file, err := os.Open("/home/moxa/Study-Note/GO/GO_Json/config")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var Addlist []AddEntry
	if err := json.Unmarshal(bytes, &Addlist); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(Addlist)
	fmt.Println(Addlist[0].Ssid)

	// data := `{"type": "time", "timezone": "Asia/Taipei", "ntp": {"enable": true, "source": "timeserver", "server": "pool.ntp.or", "gpsLongJump": true, "interval": 60}}`
	// var time TimeEntry
	// if err := json.Unmarshal([]byte(data), &time); err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(time)
	// fmt.Println(time.Ntp)
	// var ntp NtpEntry
	// if err := json.Unmarshal([]byte(data), &ntp); err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(ntp)
}
