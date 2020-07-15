package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	wpa "github.com/CPtung/wpac-go"
	"github.com/godbus/dbus/v5"
)

var (
	ifname   string
	cfile    string
	security string
	interval int32
	id       int
	ctx      context.Context
	wpacli   *wpa.WPA
)

func main() {
	var err error
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
	cfile = "/home/moxa/Study-Note/GO/GO_DBus_Wpa/wifi.conf"
	// scan()
	connect()
	// time.Sleep(5 * time.Second)
	state()
}

func state() {
	fmt.Printf("network state: %s\n", wpacli.GetInterface(ifname).State())
}

func scan() {
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
	fmt.Println(list)
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
