# gps

Using the following command to get the gps information in the container(`device_app_1`).

```[bash]
/usr/sbin/gpsd -n -b -F /var/run/gpsd.sock -P /run/gpsd/gpsd.pid
```

If the gpsd of host is running too, it will conflicts with device app.

It will change the GPS interface name if the GPSD does not release the interface during the cell_mgmt power cycle. So monitoring the kernel uevent to add/remove the interface with the `/usr/sbin/gpsdctl` can avoid the above situation.

```[go]
// monitor run monitor mode
func (s *Service) kernelUeventMonitor() {
	s.stopUEvent = make(chan struct{})
	s.logger.Printf("Monitoring UEvent kernel message to user-space...")

	conn := new(netlink.UEventConn)
	if err := conn.Connect(netlink.UdevEvent); err != nil {
		s.logger.Printf("Unable to connect to Netlink Kobject UEvent socket")
	}
	defer conn.Close()

	queue := make(chan netlink.UEvent)
	errors := make(chan error)
	quit := conn.Monitor(queue, errors, nil)

	// Signal handler to quit properly monitor mode
	go func() {
		<-s.stopUEvent
		s.logger.Printf("Exiting UEvent kernel message monitor mode...")
		quit <- struct{}{}
	}()

	// Handling message from queue
	for {
		select {
		case uevent := <-queue:
			iface := s.gps.ifaceMap[s.gps.Interface]
			if s.gps.Interface != "" && uevent.Env["DEVNAME"] == iface {
				if uevent.Action == "remove" {
					s.logger.Printf("[uevent] remove: %s", iface)
					cmd.Run(true, Gpsdctl, "remove", iface)
				} else if uevent.Action == "add" {
					s.logger.Printf("[uevent] add: %s", iface)
					cmd.Run(true, Gpsdctl, "add", iface)
				}
			}
		case <-quit:
			s.logger.Printf("Exiting UEvent monitor loop")
		}
	}
}
```
