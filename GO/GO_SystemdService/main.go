package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if err := os.MkdirAll(filepath.Dir("/etc/dlm-agent/.date/device/software/"), 0755); err != nil {
		fmt.Println(err.Error())
	}
	if err := ioutil.WriteFile("/etc/dlm-agent/.date/device/software/update.sh", []byte(UPDATE), 0755); err != nil {
		fmt.Println(err.Error())
	}
	if err := ioutil.WriteFile("/lib/systemd/system/dlm-agent-update.service", []byte(DLM_AGENT_UPDATE_SERVICE), 0755); err != nil {
		fmt.Println(err.Error())
	}
	// cmd := exec.Command("systemctl", "reset-failed", "dlm-agent-update")
	// cmd.Run()

	cmd := exec.Command("systemctl", "start", "dlm-agent-update")
	cmd.Run()
}

const DLM_AGENT_UPDATE_SERVICE = `[Unit]
Description=dlm-agent-update service

[Service]
Type=simple
ExecStart=/bin/sh /etc/dlm-agent/.date/device/software/update.sh
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=dlm-agent
[Install]
WantedBy=multi-user.target
`

const UPDATE = `#!/bin/sh
apt-get update
apt-get install dlm-agent
#rm /lib/systemd/system/dlm-agent-update.service
rm /etc/dlm-agent/.date/device/software/update.sh
`
