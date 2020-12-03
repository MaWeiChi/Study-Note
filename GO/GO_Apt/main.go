package main

import (
	"fmt"
	"strconv"

	apt "github.com/arduino/go-apt-client"
)

func main() {

	// out, err := apt.CheckForUpdates()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(string(out))
	AptListUpgradable()
	// AptList()
	// AptSearch("wpasupplicant")
	// AptUpgrade("libepub0")
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

func AptList() {
	out, err := apt.List()

	if err != nil {
		fmt.Println(err.Error())
	}
	for _, p := range out {
		size := strconv.Itoa(p.InstalledSizeKB)
		fmt.Println(p.Name + ":" + p.Version + " { " + p.Architecture + ", " + size + ", " + p.Status + ", " + p.ShortDescription + " }")
	}

}

func AptSearch(pattern string) {
	out, err := apt.Search(pattern)

	if err != nil {
		fmt.Println(err.Error())
	}
	for _, p := range out {
		size := strconv.Itoa(p.InstalledSizeKB)
		fmt.Println(p.Name + ", version: " + p.Version + " ,{ " + p.Architecture + ", " + size + ", " + p.Status + ", " + p.ShortDescription + " }")
	}

}

func AptUpgrade(pattern string) {
	upgradableList, err := apt.ListUpgradable()
	p := apt.Package{}

	for _, e := range upgradableList {
		fmt.Println(e.Name)
		if e.Name == pattern {
			p = *e
			break
		}
	}
	out, err := apt.Upgrade(&p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(out))
}
