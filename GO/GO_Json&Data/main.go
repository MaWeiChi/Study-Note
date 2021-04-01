package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func main() {

	g, _ := json.Marshal(getGeneral())
	fmt.Println(string(g))
	e, _ := json.Marshal(getEthernets())
	fmt.Println(string(e))
}

func getGeneral() DeviceInfo {
	var body = struct {
		General DeviceInfo `json:"data"`
	}{}

	var deviceInfo DeviceInfo
	url := "http://127.0.0.1:59000/api/v1/device/general"
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Errorf("Get device general failed, err: %v", err)
		return deviceInfo
	}
	req.Header.Add("mx-api-token", GetAPIToken())

	res, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Get device general failed, err: %v", err)
		return deviceInfo
	}
	defer res.Body.Close()
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("Get device general failed, err: %v", err)
		return deviceInfo
	}
	jsonData := gjson.GetBytes(response, "data").String()
	version := gjson.Get(jsonData, "thingsproVersion").String()
	j, err := sjson.Set(jsonData, "thingsProProduct", fmt.Sprintf("ThingsPro Edge V%s", strings.Split(version, "-")[0]))
	j, err = sjson.Delete(j, "thingsproVersion")
	j, err = sjson.Delete(j, "thingsproVersion")

	fmt.Print("j: ")
	fmt.Println(j)

	if err = json.Unmarshal(response, &body); err != nil {
		log.Println(err)
		errors.New("Invalid input")
		return deviceInfo
	}
	deviceInfo = body.General
	return deviceInfo

}

func getEthernets() []EthernetInfo {
	var body = struct {
		General []EthernetInfo `json:"data"`
	}{}

	var ethernetInfo []EthernetInfo
	url := "http://127.0.0.1:59000/api/v1/device/ethernets"
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Errorf("Get device general failed, err: %v", err)
		return ethernetInfo
	}
	req.Header.Add("mx-api-token", GetAPIToken())

	res, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Get device general failed, err: %v", err)
		return ethernetInfo
	}
	defer res.Body.Close()
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("Get device general failed, err: %v", err)
		return ethernetInfo
	}

	if err = json.Unmarshal(response, &body); err != nil {
		log.Println(err)
		errors.New("Invalid input")
		return ethernetInfo
	}
	ethernetInfo = body.General
	return ethernetInfo

}

func GetAPIToken() string {
	tokenFile, err := os.Open("/var/thingspro/data/mx-api-token")
	if err != nil {
		return ""
	}
	defer tokenFile.Close()
	b, err := ioutil.ReadAll(tokenFile)
	if err != nil {
		log.Printf("MOXA api token open failed, %v", err)
		return ""
	}
	return string(b)
}

type DeviceInfo struct {
	Description      string     `json:"description"`
	HostName         string     `json:"hostName"`
	ModelName        string     `json:"modelName"`
	SerialNumber     string     `json:"serialNumber"`
	FirmwareVersion  string     `json:"firmwareVersion"`
	BIOSVersion      string     `json:"biosVersion"`
	MemorySize       int        `json:"memorySize"`
	CPU              string     `json:"cpu"`
	Disk             []DiskInfo `json:"disk"`
	OSType           string     `json:"osType"`
	LastBootTime     string     `json:"lastBootTime"`
	ThingsproVersion string     `json:"thingsproVersion"`
	ThingsProProduct string     `json:"thingsProProduct"`
}

type DiskInfo struct {
	Name    string  `json:"name"`
	Mount   string  `json:"mount"`
	Device  string  `json:"device"`
	Total   uint64  `json:"total"`
	Free    uint64  `json:"free"`
	Used    uint64  `json:"used"`
	Percent float64 `json:"percent"`
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
