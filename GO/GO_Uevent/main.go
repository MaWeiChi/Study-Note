package main

import (
	"fmt"

	"github.com/pilebones/go-udev/netlink"
)

func main() {
	kernelUeventMonitor()
}

// monitor run monitor mode
func kernelUeventMonitor() {
	stopUEvent := make(chan struct{})
	fmt.Println("Monitoring UEvent kernel message to user-space...")

	conn := new(netlink.UEventConn)
	if err := conn.Connect(netlink.UdevEvent); err != nil {
		fmt.Println("Unable to connect to Netlink Kobject UEvent socket")
	}
	defer conn.Close()

	queue := make(chan netlink.UEvent)
	errors := make(chan error)
	quit := conn.Monitor(queue, errors, nil)

	// Signal handler to quit properly monitor mode
	go func() {
		<-stopUEvent
		fmt.Println("Exiting UEvent kernel message monitor mode...")
		quit <- struct{}{}
	}()

	// Handling message from queue
	for {
		select {
		case uevent := <-queue:
			fmt.Println(uevent.Action)
			fmt.Println(uevent.KObj)
			fmt.Println(uevent.Env)
			fmt.Println("")
		case <-quit:
			fmt.Println("Exiting UEvent monitor loop")
		}
	}
}
