package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/host"
)

func main() {

	if i, err := host.BootTime(); err == nil {
		fmt.Println(time.Unix(int64(i), 0).Format(time.RFC3339))
	} else {
		fmt.Println(time.Now().Format(time.RFC3339))
	}

}
