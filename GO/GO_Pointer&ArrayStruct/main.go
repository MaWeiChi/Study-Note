package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type WifiReadOnlyEntry struct {
	Regions []RegionEntry `json:"regions"`
}

type RegionEntry struct {
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
	Band24      []int  `json:"band24"`
	Band50      []int  `json:"band50"`
}

type WifiEntry struct {
	WifiReadOnlyEntry
}

type Service struct {
	wifis []WifiEntry
}

type wifisConfiguration struct {
	Wifis []WifiEntry `json:"wifis"`
}

type foo struct {
	id   int
	name string
	list []string
}

type bigFoo struct {
	Foo foo
}

func main() {
	// LoadRegions("/home/moxa/Study-Note/GO/GO_Pointer&ArrayStruct/region.json")
	// // var d1 bigFoo
	// // d1.Foo.list = append(d1.Foo.list, "1")
	// // d1.Foo.list = append(d1.Foo.list, "2")
	// // var d2 bigFoo
	// // d2.Foo = d1.Foo
	// // d2.Foo.list = []string{}
	// // fmt.Println(d1.Foo.list)

	a := make([]int, 2, 100)
	fmt.Println(a)
	a[1] = 100
	f(a)
	fmt.Println(a)
}

func f(a []int) {
	fmt.Println(a[0])
}

func LoadRegions(path string) []RegionEntry {
	configContent, _ := ioutil.ReadFile(path)
	var regions []RegionEntry
	json.Unmarshal(configContent, &regions)
	return regions
}
