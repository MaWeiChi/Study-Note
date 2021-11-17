package main

// var e1 map[string]int
// var e2 map[string]*int
// var e3 map[int]Foo
// var e4 map[int]*Foo

// // type ApServerList struct {
// // 	State   string                   `json:"state"`
// // 	List    map[string]ApServerEntry `json:"list"`
// // 	IP      string                   `json:"ip"`
// // 	Netmask string                   `json:"netmask"`
// // 	Count   int                      `json:"count"`
// // }

// // type ApServerEntry struct {
// // 	ID       int    `json:"id"`
// // 	Ssid     string `json:"ssid"`
// // 	Bssid    string `json:"bssid"`
// // 	PSK      string `json:"psk"`
// // 	Signal   int64  `json:"signal"`
// // 	Band     string `json:"band"`
// // 	KeyMgmt  string `json:"keyMgmt"`
// // 	Current  bool   `json:"current"`
// // 	Priority int    `json:"priority"`
// // }

// type Foo struct {
// 	ID   int
// 	Name string
// }

// func main() {
// 	// 	var wifi ApServerList
// 	// 	wifi.List = make(map[string]ApServerEntry)
// 	// 	// wifi.List["s"] = ApServerEntry{}
// 	// 	sside := "moxa"
// 	// 	entry := ApServerEntry{1, "ssid", "bssid", "psk", -30, "5G", "wpa", true, 2}
// 	// 	// if _, ok := wifi.List[sside]; !ok {
// 	// 	// 	wifi.List[sside] = ApServerEntry{}
// 	// 	// }
// 	// 	wifi.List[sside] = entry
// 	// 	// wifi.List[sside] = append(wifi.List[sside], entry)
// 	// 	// wifi.List["b"] = "ball"
// 	// 	fmt.Println(wifi.List)
// 	// 	fmt.Println(len(wifi.List))
// 	// 	fmt.Println(wifi)
// 	// 	// fmt.Println(len(wifi.List["moxa"]))
// 	// 	// fmt.Println(len(wifi.List["m"]))
// 	// 	// fmt.Println(wifi.List["b"])
// 	// e1 = make(map[string]int)
// 	// e1["b"] = 3
// 	// fmt.Println(e1["b"])
// 	// if i, ok := e1["a"]; ok {
// 	// 	fmt.Println(i)

// 	// } else {
// 	// 	fmt.Println("a")
// 	// }
// 	// fmt.Println(e1["a"])
// 	// e1["a"] = 0
// 	// fmt.Println(e1["a"])
// 	e4 = make(map[int]*Foo)
// 	e4[1].ID = 1
// 	// d := e4[1].ID
// 	fmt.Println(e4)
// }

// type Task struct {
// 	Cmd  string
// 	Desc string
// }

// var taskMap map[string]*Task

// // {
// // 	"showDir": {
// // 		Cmd: "ls",
// // 	},
// // 	"showDisk": {
// // 		Cmd: "df",
// // 	},
// // }

// func main() {
// 	taskMap = make(map[string]*Task)
// 	c1 := "show dirs"
// 	taskMap["showDir"].Desc = c1
// 	fmt.Print(taskMap["showDir"])
// 	var s string
// 	s = taskMap["showDir"].Desc
// 	fmt.Println(s)
// 	fmt.Print(s)
// }
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Country struct {
	Code string
	Name string
}

func main() {
	var store = make(map[string]*Country)
	var store2 = make(map[string]Country)

	c1 := Country{"US", "United States"}

	store["country1"] = &c1
	store2["country1"] = c1
	store2["s"] = Country{}
	if a, ok := store2["country1"]; ok {
		a.Code = "s"
	}
	if _, ok := store2["c"]; !ok {
		store2["c"] = Country{}
	}
	if _, ok := store2["c"]; ok {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}
	store2["c"] = c1
	c2 := Country{"TW", "Taiwan"}
	store2["c"] = c2
	fmt.Println(store["country1"].Name)  // prints "United States"
	fmt.Println(store2["country1"].Name) // prints "United States"
	fmt.Println(store2["country1"].Code)

	configJSONS, _ := json.MarshalIndent(store2, "", "  ")
	WriteFile("/home/moxa/Study-Note/GO/GO_Map/s", configJSONS, 0644)

	var data = map[string]map[string][]string{}

	data["a"] = map[string][]string{}
	data["b"] = make(map[string][]string)
	data["c"] = make(map[string][]string)

	data["a"]["w"] = []string{"x"}
	data["b"]["w"] = []string{"x"}
	data["c"]["w"] = []string{"x"}
	fmt.Println(data)
}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	tmpfile, err := ioutil.TempFile(filepath.Dir(filename), "tmp-")
	if err != nil {
		log.Print(err)
		return err
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if err := tmpfile.Chmod(perm); err != nil {
		log.Print(err)
		return err
	}
	if _, err := tmpfile.Write(data); err != nil {
		log.Print(err)
		return err
	}
	if err := tmpfile.Sync(); err != nil {
		log.Print(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Print(err)
		return err
	}

	if err := os.Rename(tmpfile.Name(), filename); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
