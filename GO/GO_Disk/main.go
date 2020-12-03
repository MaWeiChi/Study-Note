package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/shirou/gopsutil/disk"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(updateDisksInfo())
}

var mutex sync.Mutex

type DiskInfo struct {
	Name    string  `json:"name"`
	Mount   string  `json:"mount"`
	Device  string  `json:"device"`
	Total   uint64  `json:"total"`
	Free    uint64  `json:"free"`
	Used    uint64  `json:"used"`
	Percent float64 `json:"percent"`
}

func updateDisksInfo() []DiskInfo {
	mutex.Lock()
	defer mutex.Unlock()
	diskParts, err := disk.Partitions(true)
	if err != nil {
		return []DiskInfo{}
	}
	disks := []DiskInfo{}
	for _, part := range diskParts {
		// cmd.Logger.Println("mountpoint: ", part.Mountpoint)
		if part.Mountpoint == "/" {
			u, e := disk.Usage(part.Mountpoint)
			if e != nil {
				continue
			}
			logrus.Printf("Disk Usage: %v\n", u)
			d := DiskInfo{
				Name:    "System",
				Mount:   part.Mountpoint,
				Device:  "/dev/root",
				Total:   u.Total,
				Free:    u.Free,
				Used:    u.Used,
				Percent: u.UsedPercent,
			}
			disks = append(disks, d)
		} else if strings.HasPrefix(part.Mountpoint, "/host/media/") {
			// sd-mmcblkXpY -> SD (X + 1) - Y
			// label ...
			name := filepath.Base(part.Mountpoint)
			if strings.HasPrefix(name, "sd-mmcblk") && len(name) >= 12 {
				name = fmt.Sprintf("SD%s-%s", string(name[9]+1), name[11:])
			}
			u, e := disk.Usage(part.Mountpoint)
			if e != nil {
				logrus.Printf("get usage: %s\n", e.Error())
				continue
			}
			logrus.Printf("Disk Usage: %v\n", u)
			d := DiskInfo{
				Name:    name,
				Mount:   part.Mountpoint[5:], // eliminate /host
				Device:  part.Device,
				Total:   u.Total,
				Free:    u.Free,
				Used:    u.Used,
				Percent: u.UsedPercent,
			}
			disks = append(disks, d)
		}
	}
	fmt.Println("====================")
	fmt.Println("")
	fmt.Println(disks[0].Name)
	fmt.Println(disks[0].Mount)
	fmt.Println(disks[0].Device)
	fmt.Println(disks[0].Total)
	fmt.Println(disks[0].Free)
	fmt.Println(disks[0].Used)
	fmt.Println(disks[0].Percent)
	fmt.Println("")
	fmt.Println("====================")

	return disks
}
