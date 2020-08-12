package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Array struct {
	Array []string
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Person      *Person   `json:"person"`
	CreatedDate time.Time `json:"created_date"`
}

func main() {
	wifis := []WifiEntryWithNetworkLayer{}

	var wifiEntryWithNetworkLayer WifiEntryWithNetworkLayer
	var Ap1 ApServerEntry
	Ap1.Uuid = "34631c38-c4b8-11ea-87d0-0242ac130003"
	Ap1.Ssid = "moxa-test-ap-1"
	Ap1.Bssid = "10:c3:7b:c6:47:f8"
	Ap1.Security.Mode = "wpa2-personal"
	Ap1.Security.Encryption = "aes"
	Ap1.Security.Password = "!2345678"
	Ap1.Band = "Band50"
	Ap1.Signal = -17
	var Ap2 ApServerEntry
	Ap2.Uuid = "34631e54-c4b8-11ea-87d0-0242ac130003"
	Ap2.Ssid = "moxa-ap-2"
	Ap2.Security.Mode = "wpa2-personal"
	Ap2.Security.Encryption = "aes"
	Ap2.Security.Password = "!2345678"
	Ap2.Band = "Band24"
	// wifiEntryWithNetworkLayer.Enable = new(bool)
	wifiEntryWithNetworkLayer.Type = "wifi"
	wifiEntryWithNetworkLayer.Id = 1
	wifiEntryWithNetworkLayer.Name = "wlan0"
	wifiEntryWithNetworkLayer.Available = true
	wifiEntryWithNetworkLayer.Mode = "client"
	wifiEntryWithNetworkLayer.Status = true
	wifiEntryWithNetworkLayer.Capabilities.Mode = []string{"ap", "client"}
	wifiEntryWithNetworkLayer.Capabilities.Band = []string{"Band24", "Band50"}
	wifiEntryWithNetworkLayer.Client.State = "connected"
	wifiEntryWithNetworkLayer.Client.CurrentAp = "34631c38-c4b8-11ea-87d0-0242ac130003"
	wifiEntryWithNetworkLayer.Client.Priority = []string{"34631c38-c4b8-11ea-87d0-0242ac130003", "34631e54-c4b8-11ea-87d0-0242ac130003"}
	wifiEntryWithNetworkLayer.Client.Networks = append(wifiEntryWithNetworkLayer.Client.Networks, Ap1)
	wifiEntryWithNetworkLayer.Client.Networks = append(wifiEntryWithNetworkLayer.Client.Networks, Ap2)
	wifiEntryWithNetworkLayer.Client.Networks[0] = Ap1
	wifiEntryWithNetworkLayer.Client.Networks[1] = Ap2
	wifiEntryWithNetworkLayer.Client.EnableDhcp = false
	wifiEntryWithNetworkLayer.Client.Ip = "192.168.3.100"
	wifiEntryWithNetworkLayer.Client.Netmask = "255.255.252.0"
	wifiEntryWithNetworkLayer.Client.Checkalive.Enable = true
	wifiEntryWithNetworkLayer.Client.Checkalive.TargetHost = "8.8.8.8"
	wifiEntryWithNetworkLayer.Client.Checkalive.IntervalSec = 60
	wifis = append(wifis, wifiEntryWithNetworkLayer)

	router := gin.Default()
	router.GET("", func(c *gin.Context) {

		c.JSON(200, gin.H{"data": wifis, "count": len(wifis)})
	})

	router.Run()

}

func GetMockData(c *gin.Context) {
	wifis := []WifiEntryWithNetworkLayer{}

	var wifiEntryWithNetworkLayer WifiEntryWithNetworkLayer
	var Ap1 ApServerEntry
	Ap1.Uuid = "34631c38-c4b8-11ea-87d0-0242ac130003"
	Ap1.Ssid = "moxa-test-ap-1"
	Ap1.Bssid = "10:c3:7b:c6:47:f8"
	Ap1.Security.Mode = "wpa2-personal"
	Ap1.Security.Encryption = "aes"
	Ap1.Security.Password = "!2345678"
	Ap1.Band = "Band50"
	Ap1.Signal = -17
	var Ap2 ApServerEntry
	Ap2.Uuid = "34631e54-c4b8-11ea-87d0-0242ac130003"
	Ap2.Ssid = "moxa-ap-2"
	Ap2.Security.Mode = "wpa2-personal"
	Ap2.Security.Encryption = "aes"
	Ap2.Security.Password = "!2345678"
	Ap2.Band = "Band24"

	wifiEntryWithNetworkLayer.Type = "wifi"
	wifiEntryWithNetworkLayer.Id = 1
	wifiEntryWithNetworkLayer.Name = "wlan0"
	wifiEntryWithNetworkLayer.Available = true
	wifiEntryWithNetworkLayer.Mode = "client"
	wifiEntryWithNetworkLayer.Status = true
	wifiEntryWithNetworkLayer.Capabilities.Mode = []string{"ap", "client"}
	wifiEntryWithNetworkLayer.Capabilities.Band = []string{"Band24", "Band50"}
	wifiEntryWithNetworkLayer.Client.State = "connected"
	wifiEntryWithNetworkLayer.Client.CurrentAp = "34631c38-c4b8-11ea-87d0-0242ac130003"
	wifiEntryWithNetworkLayer.Client.Priority = []string{"34631c38-c4b8-11ea-87d0-0242ac130003", "34631e54-c4b8-11ea-87d0-0242ac130003"}
	wifiEntryWithNetworkLayer.Client.Networks = append(wifiEntryWithNetworkLayer.Client.Networks, Ap1)
	wifiEntryWithNetworkLayer.Client.Networks = append(wifiEntryWithNetworkLayer.Client.Networks, Ap2)
	wifiEntryWithNetworkLayer.Client.Networks[0] = Ap1
	wifiEntryWithNetworkLayer.Client.Networks[1] = Ap2
	wifiEntryWithNetworkLayer.Client.EnableDhcp = false
	wifiEntryWithNetworkLayer.Client.Ip = "192.168.3.100"
	wifiEntryWithNetworkLayer.Client.Netmask = "255.255.252.0"
	wifiEntryWithNetworkLayer.Client.Checkalive.Enable = true
	wifiEntryWithNetworkLayer.Client.Checkalive.TargetHost = "8.8.8.8"
	wifiEntryWithNetworkLayer.Client.Checkalive.IntervalSec = 60
	wifis = append(wifis, wifiEntryWithNetworkLayer)
	c.JSON(
		http.StatusOK,
		gin.H{
			"data":  wifis,
			"count": len(wifis),
		},
	)
}

type WifiEntryWithNetworkLayer struct {
	WifiReadOnlyEntry
	WifiReadWriteEntryWithNetworkLayer
}

type WifiReadOnlyEntry struct {
	Type         string          `json:"type"`
	Name         string          `json:"name"`
	Available    bool            `json:"available"`
	Status       bool            `json:"status"`
	Capabilities CapabilityEntry `json:"capabilities"`
	Regions      []RegionEntry   `json:"regions"`
}

type WifiReadWriteEntryWithNetworkLayer struct {
	Id     int                             `json:"id,omitempty" validate:"omitempty,min=1"`
	Enable *bool                           `json:"enable,omitempty"`
	Mode   string                          `json:"mode,omitempty" validate:"omitempty,oneof=ap client"`
	Ap     ApEntry                         `json:"ap,omitempty"`
	Client WifiClientEntryWithNetworkLayer `json:"client,omitempty"`
}

type WifiClientEntryWithNetworkLayer struct {
	WifiClientEntry
	EnableDhcp bool            `json:"enableDhcp"`
	Ip         string          `json:"ip"`
	Netmask    string          `json:"netmask"`
	Checkalive CheckaliveEntry `json:"checkalive"`
}

type CheckaliveEntry struct {
	IntervalSec int    `json:"intervalSec,omitempty" validate:"required,min=10,max=86400"`
	Enable      bool   `json:"enable"`
	TargetHost  string `json:"targetHost,omitempty" validate:"required,min=2,max=253,ipv4|fqdn"`
}

type ClientEntry struct {
	Mac string `json:"mac"`
}

type SecurityEntry struct {
	Mode       string `json:"mode" validate:"omitempty,eq=wpa2"`      // only wpa2 now
	Encryption string `json:"encryption" validate:"omitempty,eq=aes"` // only aes now
	Password   string `json:"password" validate:"omitempty,min=8,max=63,printascii"`
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

type WifiClientEntry struct {
	State     string          `json:"state"`
	CurrentAp string          `json:"currentAp"`
	Priority  []string        `json:"priority"`
	Networks  []ApServerEntry `json:"networks"`
}

type ApServerEntry struct {
	Uuid     string        `json:"uuid"`
	Ssid     string        `json:"ssid"`
	Bssid    string        `json:"bssid"`
	Security SecurityEntry `json:"security"`
	Signal   int           `json:"signal"`
	Band     string        `json:"band"`
}

type ScanList struct {
	ScanList []ScanEntry `json:"scanList"`
}

type ScanEntry struct {
	Ssid         string `json:"ssid"`
	Bssid        string `json:"Bssid"`
	Signal       int    `json:"Signal"`
	Band         string `json:"band"`
	KeyMgmt      string `json:"KeyMgmt"`
	bssidArray   []string
	signalArray  []int
	bandArray    []int
	keyMgmtArray []string
}
