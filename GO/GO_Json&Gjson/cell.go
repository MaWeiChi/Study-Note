package main

type EthernetInfo struct {
	DisplayName string   `json:"displayName"`
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Enable      bool     `json:"enable"`
	Status      string   `json:"status"`
	IP          string   `json:"ip"`
	Netmask     string   `json:"netmask"`
	Subnet      string   `json:"subnet"`
	Broadcast   string   `json:"broadcast"`
	Gateway     string   `json:"gateway"`
	DNS         []string `json:"dns"`
	MAC         string   `json:"mac"`
	Speed       string   `json:"speed"`
}

type CellularInfo struct {
	EthernetInfo
	Strength string  `json:"strength"`
	RSSI     float32 `json:"rssi"`
	RSRP     int     `json:"rsrp"`
	Level    int     `json:"level"`
}

func convertCellularSignalLevel(level int) string {
	switch level {
	case 0:
		return "No Signal"
	case 1:
		return "Very Poor"
	case 2:
		return "Poor"
	case 3:
		return "Fair"
	case 4:
		return "Good"
	case 5:
		return "Excellent"
	}
	return "Unknown"
}
