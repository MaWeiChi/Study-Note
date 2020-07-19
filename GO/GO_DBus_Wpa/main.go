package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/CPtung/wpac-go"
	wpa "github.com/CPtung/wpac-go"
	"github.com/godbus/dbus/v5"
)

var (
	ifname        string
	cfile         string
	security      string
	interval      int32
	id            int
	ctx           context.Context
	wpacli        *wpa.WPA
	priorityArray map[string]int
)

type uuidInformation struct {
	netwroksId          int
	networksPriority    int
	clientNetworksIndex int
}
type WifiEntry struct {
	WifiReadOnlyEntry
	WifiReadWriteEntry
}

type WifiReadWriteEntry struct {
	Id     int             `json:"id,omitempty" validate:"omitempty,min=1"`
	Enable bool            `json:"enable"`
	Mode   string          `json:"mode,omitempty" validate:"omitempty,oneof=ap client"`
	Client WifiClientEntry `json:"client,omitempty"`
}

type WifiReadOnlyEntry struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
	Status    bool   `json:"status"`
}

type WifiClientEntry struct {
	State      string          `json:"state"`
	CurrentAp  string          `json:"currentAp"`
	Priorities []string        `json:"priorities"`
	Networks   []ApServerEntry `json:"networks"`
}

type ApServerEntry struct {
	Uuid     string        `json:"uuid"`
	Ssid     string        `json:"ssid"`
	Bssid    string        `json:"bssid"`
	Security SecurityEntry `json:"security"`
	Signal   int16         `json:"signal"`
	Band     string        `json:"band"`
}

type ScanListEntry struct {
	Networks []ScanEntry `json:"networks"`
	apExits  map[string]int16
}

type ScanEntry struct {
	Ssid     string        `json:"ssid"`
	Security SecurityEntry `json:"security"`
	Signal   int16         `json:"signal"`
	Band     string        `json:"band"`
}

type SecurityEntry struct {
	Mode       string `json:"mode" validate:"omitempty,eq=wpa2"`      // only wpa2 now
	Encryption string `json:"encryption" validate:"omitempty,eq=aes"` // only aes now
	Password   string `json:"password,omitempty" validate:"omitempty,min=8,max=63,printascii"`
	Support    *bool  `json:"support,omitempty"`
}

func state() {
	fmt.Printf("network state: %s\n", wpacli.GetInterface(ifname).State())
}

func scan() {
	var ScanList ScanListEntry
	ScanList.apExits = make(map[string]int16)
	if interval == 0 {

	} else if interval > 0 {
		err := wpacli.GetInterface(ifname).SetScanInterval(interval)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	list, err := wpacli.GetInterface(ifname).AutoScan()
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, bss := range list {
		if ScanList.apExits[bss.SSID] == 0 {
			var ScanedAP ScanEntry
			ScanedAP.Ssid = bss.SSID
			ScanedAP.Band = convertTerms("", "", bss.Frequency)
			assignWpaStruct(&ScanedAP.Security, bss.WPA, bss.WPA2)
			ScanList.Networks = append(ScanList.Networks, ScanedAP)
			ScanList.apExits[bss.SSID] = bss.Signal
		} else if ScanList.apExits[bss.SSID] < bss.Signal {
			ScanList.apExits[bss.SSID] = bss.Signal
		}
	}
	for i := 0; i < len(ScanList.Networks); i++ {
		ScanList.Networks[i].Signal = ScanList.apExits[ScanList.Networks[i].Ssid]
	}
	j, _ := json.Marshal(ScanList.Networks)
	fmt.Println(string(j))
}

// func assignWpaStruct(s *SecurityEntry, w1 *wpac.BSSWPA, w2 *wpac.BSSWPA2) {
// 	if w2 != nil {
// 		s.Mode = convertTerms(strings.Split(w2.KeyMgmt[0], "-")[1], "wpa2", 0)
// 		s.Encryption = convertTerms(w2.Group, "", 0)
// 		if s.Mode == "wpa2-personal" {
// 			s.Support = true
// 		}
// 	} else if w1 != nil {
// 		s.Mode = convertTerms(strings.Split(w1.KeyMgmt[0], "-")[1], "wpa", 0)
// 		s.Encryption = convertTerms(w1.Group, "", 0)
// 		s.Support = false
// 	} else {
// 		s.Mode = "none"
// 		s.Encryption = "none"
// 		s.Support = true
// 	}
// }

func assignWpaStruct(s *SecurityEntry, w1 *wpac.BSSWPA, w2 *wpac.BSSWPA2) {
	s.Support = new(bool)
	if w2 != nil {
		s.Mode = convertTerms(strings.Split(w2.KeyMgmt[0], "-")[1], "wpa2", 0)
		s.Encryption = convertTerms(w2.Group, "", 0)
		if s.Mode == "wpa2-personal" {
			*s.Support = true
		} else {
			*s.Support = false
		}
	} else if w1 != nil {
		s.Mode = convertTerms(strings.Split(w1.KeyMgmt[0], "-")[1], "wpa", 0)
		s.Encryption = convertTerms(w1.Group, "", 0)
		*s.Support = false
	} else {
		s.Mode = "none"
		s.Encryption = "none"
		*s.Support = true
	}
}

func convertTerms(s, t string, i uint16) string {
	if s == "psk" {
		return t + "-personal"
	} else if s == "eap" {
		return t + "-enterprise"
	} else if s == "ccmp" {
		return "aes"
	} else if s == "tkip" {
		return s
	} else if i > 3000 {
		return "band50"
	} else if 0 < i {
		return "band24"
	}
	return t + s
}
func connect() {
	var (
		bss    wpa.WPABSS
		config map[string]dbus.Variant
	)

	bss = wpa.WPABSS{}
	if err := loadConfig(cfile, &bss); err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(bss)
	security = "wpa2"
	switch security {
	case "none":
		config = wpa.WPAConfig().GetWPANone(bss)
	case "wpa":
		config = wpa.WPAConfig().GetWPA2(bss)
	case "wpa2":
		config = wpa.WPAConfig().GetWPA2(bss)
	}
	network, err := wpacli.GetInterface(ifname).AddNetwork(config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := wpacli.GetInterface(ifname).SelectNetwork(network.ID); err != nil {
		log.Fatalf(err.Error())
	}
}

func disconnect() {
	err := wpacli.GetInterface(ifname).Disconnect()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func loadConfig(path string, bss *wpa.WPABSS) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	re := regexp.MustCompile(`\s*(.*)=\"*([\w\-_.@!~:]*)\"*`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) >= 3 {
			switch matches[1] {
			case "ssid":
				bss.SSID = matches[2]
			case "bssid":
				bss.BSSID = matches[2]
			case "psk":
				bss.PSK = matches[2]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func loadConfig2() ([]byte, error) {
	file, err := os.Open("/home/erik/Study-Note/GO/GO_DBus_Wpa/1.config")
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func addNetwork(bssid, ssid, psk string, priority int) {
	var (
		config map[string]dbus.Variant
	)

	bss := wpa.WPABSS{
		SSID:     ssid,
		PSK:      psk,
		Priority: priority,
	}
	if bssid != "" {
		bss.BSSID = bssid
	}
	security = "wpa2"
	switch security {
	case "none":
		config = wpa.WPAConfig().GetWPANone(bss)
	case "wpa":
		config = wpa.WPAConfig().GetWPA2(bss)
	case "wpa2":
		config = wpa.WPAConfig().GetWPA2(bss)
	}
	fmt.Println(config)
	N, err := wpacli.GetInterface(ifname).AddNetwork(config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	wpacli.GetInterface(ifname).SetNetworkEnabled(N.ID, true)
}

// func addNetworkFromEntry(networks ApServerEntry, Priority int) {

// 	var (
// 		bss    wpa.WPABSS
// 		config map[string]dbus.Variant
// 	)

// 	bss = wpa.WPABSS{}
// 	bss.SSID = networks.Ssid
// 	bss.PSK = networks.Security.Password
// 	bss.Priority = Priority
// 	security = "wpa2"
// 	switch security {
// 	case "none":
// 		config = wpa.WPAConfig().GetWPANone(bss)
// 	case "wpa":
// 		config = wpa.WPAConfig().GetWPA2(bss)
// 	case "wpa2":
// 		config = wpa.WPAConfig().GetWPA2(bss)
// 	}
// 	fmt.Println(config)
// 	N, err := wpacli.GetInterface(ifname).AddNetwork(config)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}
// 	wpacli.GetInterface(ifname).SetNetworkEnabled(N.ID, true)
// 	fmt.Println(N)
// 	fmt.Println(N.Priority)
// }

func removeNetwork() {
	err := wpacli.GetInterface(ifname).RemoveAllNetwork()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func setNetwork(networks ApServerEntry) {
	var (
		bss    wpa.WPABSS
		config map[string]dbus.Variant
	)
	bss = wpa.WPABSS{}
	bss.SSID = networks.Ssid
	bss.PSK = networks.Security.Password
	bss.Priority = 1
	bss = wpa.WPABSS{}

	switch security {
	case "none":
		config = wpa.WPAConfig().GetWPANone(bss)
	case "wpa":
		config = wpa.WPAConfig().GetWPA2(bss)
	case "wpa2":
		config = wpa.WPAConfig().GetWPA2(bss)
	}
	fmt.Println(config)
	err := wpacli.GetInterface(ifname).SetNetwork(id, config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func setPriority(priorities []string) {
	priorityArray = make(map[string]int)
	for i := 0; i < len(priorities); i++ {
		priorityArray[priorities[i]] = len(priorities) - i
	}
}

func main() {
	var err error
	// var WifiInterface []WifiEntry

	// data, err := loadConfig2()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// if err := json.Unmarshal(data, &WifiInterface); err != nil {
	// 	fmt.Println(err.Error())
	// }

	if wpacli, err = wpa.NewWPA(context.TODO()); err != nil {
		log.Fatalf(err.Error())
	}
	defer wpacli.Close()
	ifname = "wlp4s0"
	if err := wpacli.InitInterface(ifname); err != nil {
		log.Fatalf(err.Error())
	} else {

		fmt.Println("InitInterface successfully")
	}

	// // if err := wpacli.GetExistInterface(ifname); err != nil {
	// // 	log.Fatalf(err.Error())
	// // } else {

	// // 	fmt.Println("GetInterface successfully")
	// // }
	// removeNetwork()
	// fmt.Println(len(WifiInterface[0].Client.Priorities))
	// setPriority(WifiInterface[0].Client.Priorities)

	// // for _, network := range WifiInterface[0].Client.Networks {
	// // 	addNetwork(network.Ssid, network.Security.Password, priorityArray[network.Uuid])
	// // }
	// addNetwork("", "Zenfone6", "12345678", 6)
	// wpacli.GetInterface(ifname).SelectNetwork(0)
	// addNetwork("Eric", "24690885", 1)
	// addNetwork("Eric_5G", "24690885", 2)

	// // wpacli.GetInterface(ifname).SetNetworkEnabled(3, true)
	// // cfile = "/home/moxa/Study-Note/GO/GO_DBus_Wpa/wifi.conf"
	scan()
	// // connect()
	// // time.Sleep(5 * time.Second)
	// state()

	var a []ApServerEntry
	var b []ApServerEntry
	fmt.Println(NetworksEqual(a, b))

}

func NetworksEqual(a, b []ApServerEntry) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Uuid != b[i].Uuid {
			return false
		}
		if v.Security.Password != b[i].Security.Password {
			return false
		}
	}
	return true
}
