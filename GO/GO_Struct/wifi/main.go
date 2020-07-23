package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type WifiEntryWithNetworkLayer struct {
	*WifiReadOnlyEntry
	*WifiReadWriteEntry
	Client WifiClientEntryWithNetworkLayer `json:"client,omitempty"`
	logger *log.Logger
}

type WifiEntry struct {
	WifiReadOnlyEntry
	WifiReadWriteEntry
}

// type WifiReadWriteEntryWithNetworkLayer struct {
// 	Id     int                             `json:"id,omitempty" validate:"omitempty,min=1"`
// 	Enable bool                            `json:"enable"`
// 	Mode   string                          `json:"mode,omitempty" validate:"omitempty,oneof=ap client"`
// 	Ap     ApEntry                         `json:"ap,omitempty"`
// 	Client WifiClientEntryWithNetworkLayer `json:"client,omitempty"`
// }

type WifiClientEntryWithNetworkLayer struct {
	*WifiClientEntry
	IpsettingEntry
	Checkalive CheckaliveEntry `json:"checkalive"`
}

type IpsettingEntry struct {
	EnableDhcp bool   `json:"enableDhcp"`
	Ip         string `json:"ip"`
	Netmask    string `json:"netmask"`
	ethernetId int
}

type CheckaliveEntry struct {
	IntervalSec int    `json:"intervalSec,omitempty" validate:"required,min=10,max=86400"`
	Enable      bool   `json:"enable"`
	TargetHost  string `json:"targetHost,omitempty" validate:"required,min=2,max=253,ipv4|fqdn"`
}

type NetworkLayerEntry struct {
	Ipsetting  IpsettingEntry
	Checkalive CheckaliveEntry
	bs         *baseService
}

type baseService struct {
	logger *log.Logger
}

type WifiReadOnlyEntry struct {
	Type         string          `json:"type"`
	Name         string          `json:"name"`
	Available    bool            `json:"available"`
	Status       bool            `json:"status"`
	Capabilities CapabilityEntry `json:"capabilities"`
	Regions      []RegionEntry   `json:"regions"`
	Interface    string          `json:"-"`
	Phy          string          `json:"-"`
	ctrl         WifiCtrl        `json:"-"`
	clientctrl   WifiClientCtrl  `json:"-"`
}

type CapabilityEntry struct {
	Mode []string `json:"mode"`
	Band []string `json:"band"`
}

type RegionEntry struct {
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
	Band24      []int  `json:"band24"`
	Band50      []int  `json:"band50"`
}

type WifiReadWriteEntry struct {
	Id     int             `json:"id,omitempty" validate:"omitempty,min=1"`
	Enable bool            `json:"enable"`
	Mode   string          `json:"mode,omitempty" validate:"omitempty,oneof=ap client"` // only ap mode now
	Ap     ApEntry         `json:"ap,omitempty"`
	Client WifiClientEntry `json:"client,omitempty"`
}

type ClientEntry struct {
	Mac string `json:"mac"`
}

type SecurityEntry struct {
	Mode       string `json:"mode" validate:"omitempty,eq=wpa2"`      // only wpa2 now
	Encryption string `json:"encryption" validate:"omitempty,eq=aes"` // only aes now
	Password   string `json:"password" validate:"omitempty,min=8,max=63,printascii"`
	Support    *bool  `json:"support,omitempty"`
}

type ApEntry struct {
	Band          string        `json:"band" validate:"omitempty,oneof=band24 band50"`
	BroadcastSsid bool          `json:"broadcastSsid"`
	Ssid          string        `json:"ssid" validate:"omitempty,min=1,max=32,alphanum|containsany=-_"`
	Channel       int           `json:"channel"`
	Region        string        `json:"region"`
	Security      SecurityEntry `json:"security"`
	Clients       []ClientEntry `json:"clients"`
}

type WifiApImp struct {
	logger      *log.Logger
	hostapdConf string
	Interface   string
}

type WifiClientEntry struct {
	ConnectState  string          `json:"connectState"`
	CurrentAp     string          `json:"currentAp"`
	Priority      []string        `json:"priority"`
	Networks      []ApServerEntry `json:"networks"`
	NetworksAlter bool            `json:"-"`
}

type ApServerEntry struct {
	Uuid      string        `json:"uuid"`
	Ssid      string        `json:"ssid"`
	Bssid     string        `json:"bssid"`
	LockBssid bool          `json:"-"`
	Security  SecurityEntry `json:"security"`
	Signal    int16         `json:"signal"`
	Band      string        `json:"band"`
	wpaId     int
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

type WifiClientImp struct {
	logger            *log.Logger
	ctx               context.Context
	close             context.CancelFunc
	connectState      string
	currentAp         string
	interval          int32
	security          string
	priority          []string
	wpaIdMapUuid      map[int]string
	uuidMapArrayIndex map[string]int
}

type WifiCtrl interface {
	Init()
	SetInterface(_interface string)
	Update(jsonConfig string) error
	Enable() error
	Disable() error
	GetClients() []string
}

type WifiClientCtrl interface {
	Init()
	Update(*WifiClientEntry) error
	Disable()
	Scan() ([]ScanEntry, error)
	ClientGetState(*WifiClientEntry)
}

func main() {
	wifi := WifiEntry{WifiReadOnlyEntry{Type: "wifi", Name: "wlan0", Available: true, Status: true}, WifiReadWriteEntry{}}
	fmt.Println(wifi)
	wifi.Ap.Ssid = "connected"
	wifi.Client.ConnectState = "connected"
	var wifi2 WifiEntry
	wifi2.Type = "wifi"
	wifi2.Name = "wlan2"
	wifi2.Available = true
	wifi2.Status = true
	// fmt.Println(wifi2)
	var wifiwith3 WifiEntryWithNetworkLayer
	file, err := os.Open("/home/moxa/Study-Note/GO/GO_Struct/wifi/1")
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	wifiwith3.WifiReadOnlyEntry = &wifi.WifiReadOnlyEntry
	wifiwith3.WifiReadWriteEntry = &wifi.WifiReadWriteEntry
	wifiwith3.Client.WifiClientEntry = &wifi.Client
	err = json.Unmarshal(data, &wifiwith3)
	if err != nil {
		fmt.Println(err.Error())
	}
	j3, _ := json.Marshal(wifiwith3)

	fmt.Println(string(j3))
	wifiwith3.WifiReadOnlyEntry = &wifi.WifiReadOnlyEntry
	wifiwith3.WifiReadWriteEntry = &wifi.WifiReadWriteEntry
	wifiwith3.Client.WifiClientEntry = &wifi.Client
	wifiwith3.Client.Ip = "8.8.8.8"
	fmt.Println(wifiwith3.Name)

	j, _ := json.Marshal(wifi)
	j3, _ = json.Marshal(wifiwith3)
	fmt.Println(string(j))
	fmt.Println(string(j3))
}
