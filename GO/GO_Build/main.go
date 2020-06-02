package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

type Information struct {
	Models  string `json:"models"`
	Version string `json:"version"`
}

func main() {
	info := getModuleAndSoftwareInfo()

	configJSONS, _ := json.MarshalIndent(info, "", "  ")
	_ = ioutil.WriteFile("./info", configJSONS, 0644)
	fmt.Println(string(configJSONS))

	fmt.Println(info)
}

func getModuleAndSoftwareInfo() Information {
	var info Information
	version, _ := exec.Command("/bin/bash", "-c", `uname -mrs`).CombinedOutput()
	versions := strings.Split(string(version), "\n")[0]
	info.Models = strings.Split(versions, " ")[0]
	info.Version = strings.Split(versions, " ")[1]
	return info

}
