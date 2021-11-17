package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/net"

	"github.com/tidwall/sjson"
)

type Statistic struct {
	IfaceSet            map[string]string // { name: displayName }
	tagGroupListenerSet map[string]bool

	Logger *log.Logger
	mutex  sync.Mutex
}

// MemoryUsage ...
type MemoryUsage struct {
	Total uint64 `name:"memoryTotal"`

	// Used = Total - Free - Buffers - Cache
	Used uint64 `name:"memoryUsed"`

	Free      uint64 `name:"memoryFree"`
	Shared    uint64 `name:"memoryShared"`
	Buffers   uint64 `name:"memoryBuffers"`
	Cached    uint64 `name:"memoryCached"`
	Available uint64 `name:"memoryAvailable"`

	// Usage = Used * 100 / Total
	Usage uint64 `name:"memoryUsage"`
}

// StatisticEntry ...
type StatisticEntry struct {
	Name          string   `json:"name"`
	DisplayName   string   `json:"displayName"`
	LastResetTime uint64   `json:"lastResetTime"`
	Timestamp     int64    `json:"timestamp"`
	Total         uint64   `json:"total"`
	Transmit      uint64   `json:"transmit"`
	Receive       uint64   `json:"receive"`
	Tags          []string `json:"tags"`
}

func main() {
	// v, _ := mem.VirtualMemory()

	// // almost every return value is a struct
	// fmt.Printf("Total: %v, Free:%v, Used:%v\n", v.Total, v.Free, v.Used)

	// // convert to JSON. String() is also implemented
	// // fmt.Println(v)

	// ctotal, _ := cpu.Times(false)
	// fmt.Println(uint64(ctotal[0].Idle))
	// fmt.Println(uint64(ctotal[0].Total()))

	// clist, _ := cpu.Times(true)
	// // fmt.Println(len(c))

	// // fmt.Println(uint64(100.0 - (100.0 * c[0].Idle / c[0].Total())))
	// // fmt.Println(uint64(100.0 - (100.0 * c[1].Idle / c[1].Total())))
	// // fmt.Println(uint64(100.0 - (100.0 * c[2].Idle / c[2].Total())))
	// // fmt.Println(uint64(100.0 - (100.0 * c[7].Idle / c[7].Total())))

	// fmt.Println(uint64(clist[2].Idle))
	// fmt.Println(uint64(clist[2].Total()))

	// // fmt.Println(c[1].Total())
	// // fmt.Println(c[2].CPU)
	// // fmt.Println(c[3].CPU)
	// // fmt.Println(c[4].CPU)
	// // fmt.Println(c[5].CPU)
	// // fmt.Println(c[6].CPU)
	// // fmt.Println(c[7].CPU)

	// // fmt.Println(cpu.Counts(false))
	// // fmt.Println(c[0].String())
	// // StatisticEntry, _ := GetStatisticEntry()
	// // for _, Entry := range StatisticEntry {
	// // 	fmt.Println(Entry.Name)
	// // 	fmt.Println(Entry.Transmit)
	// // 	fmt.Println(Entry.Receive)

	// // }
	// vm, _ := mem.VirtualMemory()

	// mu := MemoryUsage{
	// 	Total:     vm.Total,
	// 	Free:      vm.Free,
	// 	Shared:    vm.Shared,
	// 	Buffers:   vm.Buffers,
	// 	Cached:    vm.Cached,
	// 	Available: vm.Available,
	// }
	// mu.Used = mu.Total - mu.Free - mu.Buffers - mu.Cached
	// mu.Usage = mu.Used * 100 / mu.Total
	// fmt.Println(mu.Total)
	// fmt.Println(mu.Usage)

	// for i := 0; i < 20; i++ {
	// 	getCPUUsage()
	// 	time.Sleep(time.Duration(10 * time.Second))
	// }
	// fmt.Println(GetDiskUsage("System"))
	fmt.Println(getTxRx())
}

// Get : get network statistic entries
func GetStatisticEntry() ([]StatisticEntry, error) {
	var err error
	rtn := []StatisticEntry{}

	// set "last reset time" as "boot time"
	var boot uint64
	boot, err = host.BootTime()
	if err != nil {
		return []StatisticEntry{}, fmt.Errorf("host.BootTime(): %#v", err)

	}

	var counterStats []net.IOCountersStat
	counterStats, err = net.IOCounters(true)
	if err != nil {
		return []StatisticEntry{}, fmt.Errorf("net.IOCounters(): %#v", err)
	}

	for _, cs := range counterStats {

		statistic := formStatisticEntry(
			cs,
			boot*1000, // in millisecond
		)
		rtn = append(rtn, statistic)
	}

	return rtn, nil
}

func formStatisticEntry(cs net.IOCountersStat, lastResetTime uint64) StatisticEntry {
	return StatisticEntry{
		Name:          cs.Name,
		LastResetTime: lastResetTime,
		Timestamp:     time.Now().UTC().UnixNano() / 1000000, // timestamp in millisecond
		Total:         cs.BytesSent + cs.BytesRecv,
		Transmit:      cs.BytesSent,
		Receive:       cs.BytesRecv,
		Tags: []string{
			strings.ToLower(cs.Name) + "NetowkrUsage",
			strings.ToLower(cs.Name) + "NetworkTx",
			strings.ToLower(cs.Name) + "NetworkRx",
		},
	}
}

// CPUUsage ...
type CPUUsage struct {
	Name    string  `json: "name"`
	User    float64 `name:"cpuUser"`
	Nice    float64 `name:"cpuNice"`
	System  float64 `name:"cpuSystem"`
	Idle    float64 `name:"cpuIdle"`
	Iowait  float64 `name:"cpuIowait"`
	Irq     float64 `name:"cpuIrq"`
	Softirq float64 `name:"cpuSoftirq"`
	Steal   float64 `name:"cpuSteal"`

	// Usage = 100 - (100 * Idle / Total)
	Usage uint64 `name:"cpuUsage"`

	Total float64
}

var leastCPUsInfo []CPUUsage

func GetCPUs() []CPUUsage {
	cpus := []CPUUsage{}
	times, _ := cpu.Times(true)

	for i := range times {
		name := fmt.Sprintf("CPU-%d", i+1)
		cu := CPUUsage{
			Name:    name,
			User:    times[i].User,
			Nice:    times[i].Nice,
			System:  times[i].System,
			Idle:    times[i].Idle,
			Iowait:  times[i].Iowait,
			Irq:     times[i].Irq,
			Softirq: times[i].Softirq,
			Steal:   times[i].Steal,
			Total:   times[i].Total(),
		}
		cpus = append(cpus, cu)
	}

	return cpus
}

func getCPUUsage() {
	details := ``
	var totalUsage uint64
	newCPUsInfo := GetCPUs()
	for i := range leastCPUsInfo {
		var Usage uint64

		Usage = uint64(100.0 - 100.0*(newCPUsInfo[i].Idle-leastCPUsInfo[i].Idle)/(newCPUsInfo[i].Total-leastCPUsInfo[i].Total))
		details, _ = sjson.Set(details, fmt.Sprintf("%d."+leastCPUsInfo[i].Name, i), Usage)
		totalUsage = totalUsage + Usage
	}

	leastCPUsInfo = make([]CPUUsage, len(GetCPUs()))
	copy(leastCPUsInfo, GetCPUs())
	fmt.Println(details)
	if len(leastCPUsInfo) > 0 {
		fmt.Println(int(int(totalUsage) / len(leastCPUsInfo)))
	}
}

func taghubTestFunc(c *gin.Context) {
	tag, _ := c.GetQuery("tags")
	if !strings.Contains(tag, "cpuUsage") && !strings.Contains(tag, "memoryUsage") {
		c.Status(http.StatusOK)
		return
	}
	time.Now().Unix()
	c.String(http.StatusOK,
		`{"data":[
			{"prvdName":"system","srcName":"status","tagName":"cpuUsage","ts":1581659603,"dataUnit":"%","dataType":"int64","dataValue":76},
			{"prvdName":"system","srcName":"status","tagName":"cpuUsage","ts":1581659603,"dataUnit":"%","dataType":"int64","dataValue":76},
			{"prvdName":"system","srcName":"status","tagName":"memoryUsage","ts":1581659603,"dataUnit":"%","dataType":"uint64","dataValue":71}]}`)
}

type DiskUsage struct {
	Used    uint64
	Free    uint64
	Percent float64
}

// GetDiskUsage ...
func GetDiskUsage(mountpoint string) (DiskUsage, error) {
	u, e := disk.Usage(mountpoint)
	if e != nil {
		// cmd.Logger.Printf("Error: disk.Usage(%s) %v\n", mountpoint, e)
		return DiskUsage{}, e
	}

	du := DiskUsage{
		Used:    u.Used,
		Free:    u.Free,
		Percent: u.UsedPercent,
	}
	// cmd.Logger.Printf("disk.Usage(%s): %v\n", mountpoint, du)

	return du, nil
}

var (
	excludeRegex = []string{
		`^lo$`,
		`^docker0$`,
		`^veth[a-z0-9@]+$`,
		`^br-[a-z0-9]+$`,
	}
)

var leastTx map[string]uint64
var leastRx map[string]uint64

func getTxRx() (tx, rx map[string]uint64) {
	tx = make(map[string]uint64)
	rx = make(map[string]uint64)
	counterStats, _ := net.IOCounters(true)

	exclude := strings.Join(excludeRegex, "|")
	for _, cs := range counterStats {
		if match, ok := regexp.MatchString(exclude, cs.Name); match || (ok != nil) {
			continue
		}
		tx[cs.Name] = cs.BytesSent
		rx[cs.Name] = cs.BytesRecv
	}

	return tx, rx
}
