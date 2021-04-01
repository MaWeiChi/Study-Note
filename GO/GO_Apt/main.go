package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	apt "github.com/arduino/go-apt-client"
)

func main() {

	// out, err := apt.CheckForUpdates()
	// list, _ := ListUpgradable()
	// fmt.Println(len(list))
	// t()
	// list, _ := getUpgradableList()
	// dlist := getUpgradableListDetial(list)
	// fmt.Println(len(list))
	// fmt.Println(dlist[0].DownloadSizeKB)
	// fmt.Println(dlist[0].InstalledSizeKB)
	// for _, d := range dlist[].Depends {
	// 	fmt.Println(d)

	// }
	// fmt.Println(dlist[1].InstalledSizeKB)
	// fmt.Println(dlist[1])

	// fmt.Println(list[0])
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(string(out))
	// AptListUpgradable()
	// AptList()
	// AptSearch("wpasupplicant")
	// AptUpgrade("libepub0")
	dryrun()

}

func AptUpdate() {
	out, err := apt.CheckForUpdates()

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(out))
}

func AptListUpgradable() {
	out, err := apt.ListUpgradable()

	if err != nil {
		fmt.Println(err.Error())
	}
	for _, p := range out {
		size := strconv.Itoa(p.InstalledSizeKB)
		fmt.Println(p.Name + ":" + p.Version + " { " + p.Architecture + ", " + size + ", " + p.Status + ", " + p.ShortDescription + " }")
	}

}

// func AptList() {
// 	out, err := apt.List()

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	for _, p := range out {
// 		size := strconv.Itoa(p.InstalledSizeKB)
// 		fmt.Println(p.Name + ":" + p.Version + " { " + p.Architecture + ", " + size + ", " + p.Status + ", " + p.ShortDescription + " }")
// 	}

// }

// func AptList() {
// 	cmd := exec.Command("apt-get", "list")

// 	out, _ := cmd.CombinedOutput()

// 	res := []*Package{}
// 	scanner := bufio.NewScanner(bytes.NewReader(out))
// 	for scanner.Scan() {
// 		data := strings.Split(scanner.Text(), "\t")
// 		size, err := strconv.Atoi(data[4])
// 		if err != nil {
// 			// Ignore error
// 			size = 0
// 		}
// 		res = append(res, &Package{
// 			Name:             data[0],
// 			Architecture:     data[1],
// 			Status:           data[2],
// 			Version:          data[3],
// 			InstalledSizeKB:  size,
// 			ShortDescription: data[5],
// 		})
// 	}
// 	return res
// }

// func AptUpgrade(pattern string) {
// 	upgradableList, err := apt.ListUpgradable()
// 	p := apt.Package{}

// 	for _, e := range upgradableList {
// 		fmt.Println(e.Name)
// 		if e.Name == pattern {
// 			p = *e
// 			break
// 		}
// 	}
// 	out, err := apt.Upgrade(&p)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(string(out))
// }

func t() {
	cmd := exec.Command("apt", "list", "--upgradable")
	grep := exec.Command("grep", "ap")

	// Get ps's stdout and attach it to grep's stdin.
	pipe, _ := cmd.StdoutPipe()
	defer pipe.Close()

	grep.Stdin = pipe

	// Run ps first.
	cmd.Start()

	// Run and get the output of grep.
	res, _ := grep.Output()

	fmt.Println(string(res))
}

func dryrun() []string {
	deps := []string{}
	cmd := exec.Command("apt", "install", "--dry-run", "nvidia-prime")
	out, err := cmd.Output()
	if err != nil {
		log.Printf("apt get upgradable list failed: %s\n,", err.Error())
		return deps
	}
	s := strings.Split(string(out), "\n")
	for i := range s {
		// if s[i] == "The following packages will be upgraded:" {
		// 	deps = strings.Split(strings.TrimSpace(s[i+1]), " ")
		// }
		if conf := strings.Split(strings.TrimSpace(s[i]), " "); conf[0] == "Conf" {
			deps = append(deps, conf[1])
			fmt.Println(conf[2][1:])
		}
	}
	fmt.Println(deps)
	return deps
}

func getUpgradableList() ([]Package, error) {
	cmd := exec.Command("apt", "list", "--upgradable")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("running apt list: %s", err)
	}
	re := regexp.MustCompile(`^([^ ]+) ([^ ]+) ([^ ]+) \[upgradable from: (.+)]`)

	res := []Package{}
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		if len(matches) == 0 {
			continue
		}
		name := strings.Split(matches[0][1], "/")[0]

		res = append(res, Package{
			Name:         name,
			Status:       "upgradable",
			Version:      matches[0][2],
			Architecture: matches[0][3],
			FromVersion:  matches[0][4],
		})
	}
	return res, nil
}

func getUpgradableListDetial(list []Package) []Package {
	res := []Package{}

	for _, p := range list {
		cmd := exec.Command("apt-cache", "--no-all-versions", "show", p.Name)
		out, _ := cmd.Output()
		re := regexp.MustCompile(`^([^\ ]+): ([^\n]+)`)
		var downloadSize int
		// var installedSize int
		var shortDescription string
		var maintainer string
		depends := []string{}
		scanner := bufio.NewScanner(bytes.NewReader(out))
		for scanner.Scan() {
			matches := re.FindAllStringSubmatch(scanner.Text(), -1)
			if len(matches) == 0 {
				continue
			}
			if matches[0][1] == "Maintainer" {
				maintainer = matches[0][2]
			}
			// if matches[0][1] == "Installed-Size" {
			// 	ssize := strings.Split(matches[0][2], " kB")[0]
			// 	installedSize, _ = strconv.ParseFloat(ssize, 64)
			// 	installedSize = installedSize * 1024

			// }
			if matches[0][1] == "Depends" {
				depends = strings.Split(matches[0][2], ", ")
			}
			if matches[0][1] == "Size" {
				// ssize := strings.Split(matches[0][2], " kB")[0]
				downloadSize, _ = strconv.Atoi(matches[0][2])
				downloadSize = downloadSize * 1024

				// 	downloadSize, _ = strconv.Atoi(ssize)
			}
			if matches[0][1] == "Description" {
				shortDescription = matches[0][2]
			}
		}
		res = append(res, Package{
			Name:         p.Name,
			Status:       "upgradable",
			Architecture: p.Architecture,
			Version:      p.Version,
			FromVersion:  p.FromVersion,
			Maintainer:   maintainer,
			// InstalledSizeKB:  installedSize,
			Depends:          depends,
			DownloadSizeKB:   downloadSize,
			ShortDescription: shortDescription,
		})

	}
	return res
}

type Package struct {
	Name           string
	Status         string
	Architecture   string
	Version        string
	FromVersion    string
	DownloadSizeKB int
	// InstalledSizeKB  int
	ShortDescription string
	Maintainer       string
	Depends          []string
}
