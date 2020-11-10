package main

type Service struct {
	layer3Set map[int]*NetworkLayerEntry
}

type NetworkLayerEntry struct {
	Ipsetting  IpsettingEntry  `json:"ipsetting,omitempty"`
	Checkalive CheckaliveEntry `json:"checkalive,omitempty"`
}

type IpsettingEntry struct {
	EnableDhcp bool     `json:"enableDhcp,omitempty"`
	Ip         string   `json:"ip,omitempty" validate:"omitempty,ipv4`
	Netmask    string   `json:"netmask,omitempty"`
	Gateway    string   `json:"gateway,omitempty" validate:"omitempty,ipv4,gateway"`
	Dns        []string `json:"dns,omitempty" validate:"omitempty,dive,omitempty,ipv4"`
	EthernetId int      `json:"-"`
}

type CheckaliveEntry struct {
	IntervalSec        int    `json:"intervalSec,omitempty" validate:"require,min=10,max=86400"`
	Enable             bool   `json:"enable"`
	TargetHost         string `json:"targetHost,omitempty" validate:"require,min=2,max=253,ipv4|fqdn"`
	Layer3Successfully bool   `json:"-"`
	chkRunning         bool
}

type WifiEntryWithNetworkLayer struct {
	*WifiReadOnlyEntry
	*WifiReadWriteEntry
	Client  WifiClientEntryWithNetworkLayer `json:"client,omitempty"`
	keepAch chan bool
}

type WifiClientEntryWithNetworkLayer struct {
	WifiClientEntry
	Ipsetting  IpsettingEntry   `json:"ipsetting,omitempty"`
	Checkalive *CheckaliveEntry `json:"checkalive,omitempty" validate:"omitempty"`
}

type WifiReadOnlyEntry struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Available   bool   `json:"available"`
	Status      bool   `json:"status"`
	Interface   string `json:"-"`
	Phy         string `json:"-"`
	lastMode    string
}

type WifiReadWriteEntry struct {
	Id     int             `json:"id,omitempty" validate:"omitempty,min=1"`
	Enable bool            `json:"enable"`
	Mode   string          `json:"mode,omitempty" validate:"omitempty,oneof=ap client"`
	Ap     ApEntry         `json:"ap,omitempty"`
	Client WifiClientEntry `json:"client,omitempty"`
}

type ApEntry struct {
	Band          string `json:"band" validate:"omitempty,oneof=band24 band50"`
	BroadcastSsid bool   `json:"broadcastSsid"`
	Ssid          string `json:"ssid" validate:"omitempty,min=1,max=32,alphanum|containsany=-_"`
	Channel       int    `json:"channel"`
	Region        string `json:"region"`
}

type WifiClientEntry struct {
	ConnectState string   `json:"connectState"`
	CurrentAp    string   `json:"currentAp"`
	Priority     []string `json:"priority"`
}

type CheckAliveEntry struct {
	gatewayMac string
	targetIP   string
	TargetHost string
	Gateway    string
	DNS        []string
	Interface  string
	Interval   int
}
