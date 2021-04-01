package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

func main() {
	request()
}

func request() {

	url := "http://10.144.48.143/api/v1/device/cellulars"
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Errorf("Get device general failed, err: %v", err)

	}
	req.Header.Add("mx-api-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MCwiVXNlcm5hbWUiOiJhZG1pbiIsIlBlcm1pc3Npb25zIjpbIlNZU19NWEFQSV9UT0tFTiJdLCJleHAiOjQ3NjYwMDQ4MTN9.nuv-27KFFDX3f5F0Dh016e_dRRpzQx5IXCZk4nzfK64")

	res, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Get device general failed, err: %v", err)
	}
	// fmt.Println(res)
	response, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(response))
	f := `{"count":1,"data":[{"available":true,"status":"disconnected","currentProfileId":0,"operatorName":"","name":"usb0","module":"u-blox TOBY-L2 series","imei":"352255062111061","iccId":"","imsi":"","pinRetryRemain":0,"mac":"02:01:05:19:00:0b","rat":"","signal":{"rat":"","csq":0,"rssi":0,"rxqual":0,"rscp":3,"ecio":0,"rsrp":3,"rsrq":0,"indicator":"","level":0},"type":"cellulars","wan":true,"displayName":"Cellular1","capabilities":{"sim":2},"id":1,"profileTimeout":120,"enable":false,"keepalive":{"intervalSec":60,"enable":true,"targetHost":"8.8.8.8","reboot":{"enable":false,"intervalMin":20}},"autoDetect":true,"profiles":[]}]}`
	var c []CellularInfo
	d := gjson.Get(f, `data`).String()
	json.Unmarshal([]byte(d), &c)
	fmt.Println(c)

	pattern := fmt.Sprintf(`%d.signal`, 0)
	// json.Unmarshal([]byte(signal), &c)
	signal := gjson.Get(d, pattern).String()
	fmt.Println(gjson.Get(signal, `level`).String())
	// fmt.Println(gjson.GetBytes(response, pattern)
	fmt.Println(signal)
	json.Unmarshal([]byte(signal), &c[0])
	b, _ := json.Marshal(c)
	fmt.Println(string(b))
	// fmt.Println(gjson.Get(string(response), pattern).String())

}

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
	Strength string
	RSSI     int `json:"rssi"`
	RSRP     int `json:"rsrp"`
}
