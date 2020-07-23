package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var modelName = "UC-8112A-LX"

func main() {
	// fmt.Println("123 \n123")
	// fmt.Println("123 \n123")

	// fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
	// fmt.Println(readPversion().Model)
	// fmt.Println(strings.Split(modelName, "-")[1])
	t := 12345678
	s2 := strconv.Itoa(t)
	fmt.Println(t)
	// fmt.Println(string([]byte)
	fmt.Println(s2 + "sd")
}

func readPversion() FileVersionInfo {
	var deviceInfo FileVersionInfo
	if b, err := ioutil.ReadFile("/etc/pversion"); err == nil {
		b := strings.TrimSpace(string(b))
		// if string(b[len(b)-1:]) == "\n" {
		// 	b = b[:len(b)-1]
		// }

		deviceInfo.Model = strings.Split(b, " ")[0]
		deviceInfo.TPEVersion = strings.Split(b, " ")[1]

	} else {
		fmt.Errorf("Failed to catch pversion : %s", err.Error())
		deviceInfo.Model = "unknown"
		deviceInfo.TPEVersion = "unknown"
	}
	return deviceInfo
}

type FileVersionInfo struct {
	Model      string `json:"model"`
	TPEVersion string `json:"tpeversion"`
}
