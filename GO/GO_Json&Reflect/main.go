package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
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
}

type CellularInfo struct {
	EthernetInfo
	Strength string  `json:"strength"`
	RSSI     float32 `json:"rssi"`
	RSRP     float32 `json:"rsrp"`
}

func main() {

	c1 := loadingCell("/home/moxa/Erik/Study-Note/GO/GO_Json&Reflect/e1")
	c2 := loadingCell("/home/moxa/Erik/Study-Note/GO/GO_Json&Reflect/e2")
	if !reflect.DeepEqual(c1, c2) {
		fmt.Println("diff")
	}
}

func loadingCell(path string) []CellularInfo {
	var cells []CellularInfo
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf(err.Error())
		return cells
	}
	if len(data) < 1 {
		return cells
	}

	if err := json.Unmarshal(data, &cells); err != nil {
		log.Printf(err.Error())
	}

	return cells
}
