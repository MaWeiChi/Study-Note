package main

import "fmt"

type Service struct {
	wifis []WifiEntry
}
type WifiEntry struct {
	WifiReadOnlyEntry
}

type WifiReadOnlyEntry struct {
	Type        string
	Name        string
	DisplayName string
}

type WifiEntryWithNetworkLayer struct {
	*WifiReadOnlyEntry
}

func main() {
	var s Service
	var w WifiEntry
	w.Type = "client"
	w.Name = "wlan0"
	w.DisplayName = "Wi-Fi1"
	s.wifis = append(s.wifis, w)
	wifis := []WifiEntryWithNetworkLayer{}
	// fmt.Print(&s.wifis[0])
	for i := 0; i < 6; i++ {
		// w1 := s.convertLayer2toLayer3(0)

		w2 := s.convertLayer2toLayer32(0)
		fmt.Print("w2: ")
		fmt.Println(&w2)
		w3 := new(WifiEntryWithNetworkLayer)
		// w1 := s.wifis[0]
		w3.WifiReadOnlyEntry = w2.WifiReadOnlyEntry
		fmt.Print("w3: ")
		fmt.Println(&w3)
		wifis = append(wifis, *w3)
		fmt.Println(&wifis[len(wifis)-1])
		fmt.Println("")
	}

}

func (s *Service) convertLayer2toLayer3(id int) WifiEntryWithNetworkLayer {
	wifi := s.wifis[id]

	var wifiWithLayer3 WifiEntryWithNetworkLayer
	wifiWithLayer3.WifiReadOnlyEntry = &wifi.WifiReadOnlyEntry

	return wifiWithLayer3
}

func (s *Service) convertLayer2toLayer32(id int) WifiEntryWithNetworkLayer {
	var wifi WifiEntryWithNetworkLayer
	wifi.WifiReadOnlyEntry = &s.wifis[id].WifiReadOnlyEntry
	return wifi
}
