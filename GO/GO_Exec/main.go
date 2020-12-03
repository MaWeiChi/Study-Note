package main

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("apt-get", "remove", "rolldice", "-y")
	// if runtime.GOOS == "windows" {
	// 	cmd = exec.Command("tasklist")
	// }
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
