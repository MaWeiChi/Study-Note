package main

import "os"

var configurationDir = "/tmp/appman-config-golden"

func main() {
	os.MkdirAll(configurationDir, 0755)
	os.RemoveAll(configurationDir)
}
