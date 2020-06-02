package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ModuleAndTPEVersion struct {
	Module     string `json:Module`
	TPEVersion string `json:TPEVesion`
}

func main() {
	var deviceInfo ModuleAndTPEVersion
	// read /pversion under /etc
	b, _ := ioutil.ReadFile("/etc/pversion")
	if string(b[len(b)-1:]) == "\n" {
		b = b[:len(b)-1]
	}
	deviceInfo.Module = strings.Split(string(b), " ")[0]
	deviceInfo.TPEVersion = strings.Split(string(b), " ")[1]
	fmt.Println(deviceInfo)
	fmt.Println(deviceInfo.Module)
	fmt.Println(deviceInfo.TPEVersion)

	// get pwd
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	// write file under pwd dir
	jsonOutPut, _ := json.MarshalIndent(deviceInfo, "", " ")
	_ = ioutil.WriteFile(dir+"/version", jsonOutPut, 0644)

	// read json file under pwd dir
	var readDeviceInfo ModuleAndTPEVersion
	jsonRead, _ := ioutil.ReadFile(dir + "/version")
	_ = json.Unmarshal(jsonRead, &readDeviceInfo)
	fmt.Println("under pwd dir")
	fmt.Println(readDeviceInfo)

	// write file under tmp dir
	dir, err = ioutil.TempDir("", "second")
	if err != nil {
		fmt.Println(err)
	}
	_ = ioutil.WriteFile(dir+"/version", jsonOutPut, 0644)

	// read json file under tmp dir
	jsonRead, _ = ioutil.ReadFile(dir + "/version")
	_ = json.Unmarshal(jsonRead, &readDeviceInfo)
	fmt.Println("under tmp dir")
	fmt.Println(readDeviceInfo)

}
