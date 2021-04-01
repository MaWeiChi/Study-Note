package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	getModelName()
}

// type Information struct {
// 	Models  string `json:"models"`
// 	Version string `json:"version"`
// }

// func main() {
// 	info := getModuleAndSoftwareInfo()

// 	configJSONS, _ := json.MarshalIndent(info, "", "  ")
// 	_ = ioutil.WriteFile("./info", configJSONS, 0644)
// 	fmt.Println(string(configJSONS))

// 	fmt.Println(info)
// }

// func getModuleAndSoftwareInfo() Information {
// 	var info Information
// 	version, _ := exec.Command("/bin/bash", "-c", `uname -mrs`).CombinedOutput()
// 	versions := strings.Split(string(version), "\n")[0]
// 	info.Models = strings.Split(versions, " ")[0]
// 	info.Version = strings.Split(versions, " ")[1]
// 	return info

// }

func getModelName() {
	ModelName := "UNKNOWN"
	kversionPath := ""
	args := ""
	var cmd *exec.Cmd
	if _, err := os.Stat("/bin/mx-ver"); err == nil {
		kversionPath = "/bin/mx-ver"
		args = "-m"
		cmd = exec.Command(kversionPath, args)
	} else {
		if _, err := os.Stat("/bin/kversion"); err == nil {
			kversionPath = "/bin/kversion"
		} else if _, err := os.Stat("/usr/bin/kversion"); err == nil {
			kversionPath = "/usr/bin/kversion"
		} else {
			logrus.Printf("no kversion\n")
			return
		}
		cmd = exec.Command(kversionPath)
	}
	out, err := cmd.CombinedOutput()

	if err != nil {
		logrus.Printf("%v\n", err)
		return
	}

	fmt.Println(strings.TrimSpace(string(out)))
	if args != "" {
		ModelName = strings.TrimSpace(string(out))
		return
	}

	re := regexp.MustCompile(`\sversion\s`)
	arr := re.Split(strings.TrimSpace(string(out)), -1)
	if len(arr) >= 2 {
		ModelName = arr[0]
	} else {
		fmt.Println("len(arr) < 2")
		fmt.Println(arr)
	}
	fmt.Println(ModelName)
}
