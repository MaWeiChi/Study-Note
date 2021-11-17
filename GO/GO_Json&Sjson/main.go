package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/tidwall/sjson"
)

func main() {
	boo()
}

func boo() {

	s, _ := sjson.Set(``, "cpu", cpuInfo)
	fmt.Println(s)
}

type CPUInfo struct {
	Name            string         `json:"name"`
	Load            int            `json:"load"`
	Deatil          map[string]int `json:"detail"`
	UpdateTimeStamp string         `json:"updateTimeStamp"`
}

// CPUUsage ...
type CPUUsage struct {
	Name string
	// User   float64
	// Nice   float64
	// System float64
	Idle float64
	// Iowait  float64
	// Irq     float64
	// Softirq float64
	// Steal   float64

	// // Usage = 100 - (100 * Idle / Total)
	// Usage uint64

	Total float64
}

var leastCPUUsage []CPUUsage

func GetCPUs() []CPUUsage {
	cpus := []CPUUsage{}
	times, _ := cpu.Times(true)

	for i := range times {
		name := fmt.Sprintf("CPU-%d", i+1)
		cu := CPUUsage{
			Name: name,
			// User:   times[i].User,
			// Nice:   times[i].Nice,
			// System: times[i].System,
			Idle: times[i].Idle,
			// Iowait:  times[i].Iowait,
			// Irq:     times[i].Irq,
			// Softirq: times[i].Softirq,
			// Steal:   times[i].Steal,
			Total: times[i].Total(),
		}
		cpus = append(cpus, cu)
	}

	return cpus
}

var cpuInfo CPUInfo

func getCPUUsage() {
	var totalUsage int
	newCPUsInfo := GetCPUs()

	for i, c := range leastCPUUsage {
		var Usage int

		Usage = int(100.0 - 100.0*(newCPUsInfo[i].Idle-leastCPUUsage[i].Idle)/(newCPUsInfo[i].Total-leastCPUUsage[i].Total))
		cpuInfo.Deatil[c.Name] = Usage
		totalUsage = totalUsage + Usage
	}

	leastCPUUsage = make([]CPUUsage, len(GetCPUs()))
	copy(leastCPUUsage, GetCPUs())

	if len(leastCPUUsage) > 0 {
		cpuInfo.Load = int(totalUsage) / len(leastCPUUsage)
	}
	cpuInfo.UpdateTimeStamp = time.Now().Format(time.RFC3339)
}
