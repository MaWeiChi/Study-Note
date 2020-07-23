package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type wifisConfiguration struct {
	Wifis     []WifiEntry               `json:"wifis"`
	Layer3Set map[int]NetworkLayerEntry `json:"layer3Set"`
}

type WifiEntry struct {
	WifiReadOnlyEntry
	WifiReadWriteEntry
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

type ApEntry struct {
	Band          string        `json:"band" validate:"omitempty,oneof=band24 band50"`
	BroadcastSsid bool          `json:"broadcastSsid"`
	Ssid          string        `json:"ssid" validate:"omitempty,min=1,max=32,alphanum|containsany=-_"`
	Channel       int           `json:"channel"`
	Region        string        `json:"region"`
	Security      SecurityEntry `json:"security"`
	Clients       []ClientEntry `json:"clients"`
}

type SecurityEntry struct {
	Mode       string `json:"mode" validate:"omitempty,eq=wpa2"`      // only wpa2 now
	Encryption string `json:"encryption" validate:"omitempty,eq=aes"` // only aes now
	Password   string `json:"password" validate:"omitempty,min=8,max=63,printascii"`
	Support    *bool  `json:"support,omitempty"`
}

type ClientEntry struct {
	Mac string `json:"mac"`
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

type NetworkLayerEntry struct {
	Ipsetting  IpsettingEntry  `json:"ipsetting"`
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

var layer3set map[int]NetworkLayerEntry

func main() {
	layer3set = make(map[int]NetworkLayerEntry)
	var w2s []WifiEntry

	var w2 WifiEntry
	var Ap1 ApServerEntry
	Ap1.Uuid = "34631c38-c4b8-11ea-87d0-0242ac130003"
	Ap1.Ssid = "moxa-test-ap-1"
	Ap1.Bssid = "10:c3:7b:c6:47:f8"
	Ap1.Security.Mode = "wpa2-personal"
	Ap1.Security.Encryption = "aes"
	Ap1.Security.Password = "!2345678"
	Ap1.Band = "band50"
	Ap1.Signal = -17
	var Ap2 ApServerEntry
	Ap2.Uuid = "34631e54-c4b8-11ea-87d0-0242ac130003"
	Ap2.Ssid = "moxa-ap-2"
	Ap2.Security.Mode = "wpa2-personal"
	Ap2.Security.Encryption = "aes"
	Ap2.Security.Password = "!2345678"
	Ap2.Band = "band24"

	w2.Type = "wifi"
	w2.Id = 1
	w2.Name = "wlan0"
	w2.Available = true
	w2.Mode = "client"
	w2.Enable = true
	w2.Status = true
	w2.Capabilities.Mode = []string{"ap", "client"}
	w2.Capabilities.Band = []string{"band24", "band50"}
	w2.Client.ConnectState = "connected"
	w2.Client.CurrentAp = "34631c38-c4b8-11ea-87d0-0242ac130003"
	w2.Client.Priority = []string{"34631c38-c4b8-11ea-87d0-0242ac130003", "34631e54-c4b8-11ea-87d0-0242ac130003"}
	w2.Client.Networks = append(w2.Client.Networks, Ap1)
	w2.Client.Networks = append(w2.Client.Networks, Ap2)
	w2.Client.Networks[0] = Ap1
	w2.Client.Networks[1] = Ap2

	w2s = append(w2s, w2)
	if _, ok := layer3set[w2.Id]; !ok {
		layer3set[w2.Id] = NetworkLayerEntry{}
	}
	var tmp NetworkLayerEntry
	tmp.Ipsetting.EnableDhcp = false
	tmp.Ipsetting.Ip = "192.168.3.100"
	tmp.Ipsetting.Netmask = "255.255.252.0"
	tmp.Checkalive.Enable = true
	tmp.Checkalive.TargetHost = "8.8.8.8"
	tmp.Checkalive.IntervalSec = 60
	layer3set[w2.Id] = tmp

	// fmt.Println(layer3set)
	var w wifisConfiguration
	w.Wifis = w2s
	w.Layer3Set = make(map[int]NetworkLayerEntry)
	w.Layer3Set = layer3set
	// for k, v := range layer3set {
	// 	w.layer3Set[k] = v
	// }
	fmt.Println(w)
	configJSONS, _ := json.MarshalIndent(w, "", "  ")
	WriteFile("/home/moxa/Study-Note/GO/GO_Map/wifi/s", configJSONS, 0644)
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
