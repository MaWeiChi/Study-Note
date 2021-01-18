package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	schMgmt := newScheduleMgmt("test")
	schMgmt.NewSchedule("test", "", schMgmt.Print1)
	schMgmt.StartSchedule()
	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// <-quit

	time.Sleep(time.Duration(20) * time.Second)
	schMgmt.UpdateSchedule("test2", "", schMgmt.Print2, true)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
