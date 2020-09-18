package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
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

type WifiEntry struct {
	Type     string     `json:"type"`
	Networks []AddEntry `json:"networks"`
}

func main() {

	file, err := os.Open("/home/moxa/Study-Note/GO/GO_Json&Gjson/config")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var wifi WifiEntry

	if err := json.Unmarshal(bytes, &wifi); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(wifi)
	fmt.Println(wifi.Networks[0].Ssid)
	fmt.Println(json.Marshal(wifi))
	gjson.Get(string(bytes), "networks").ForEach(
		func(key, value gjson.Result) bool {
			if value.Get("ssid").Exists() {
				fmt.Println(value.Get("ssid").Str)
				return false
			}
			return false
		})

	file2, err := os.Open("/home/moxa/Study-Note/GO/GO_Json&Gjson/config2")
	if err != nil {
		panic(err)
	}
	bytes2, err := ioutil.ReadAll(file2)
	if err != nil {
		panic(err)
	}
	var Add AddEntry
	if err := json.Unmarshal(bytes2, &Add); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(Add)
	fmt.Println(Add.Ssid)
	fmt.Println(json.Marshal(Add))
	fmt.Println(gjson.Get(string(bytes2), "*d").Str)
	fmt.Println(gjson.Get(string(bytes2), "p*d").Str)

	fmt.Println(gjson.Get(string(bytes2), "pas?word").Str)
	if gjson.Get(string(bytes2), "mode").Str == "" {
		fmt.Println(`if the gjson doesn't exist, it will retrun a result with "" str`)
	}
}
