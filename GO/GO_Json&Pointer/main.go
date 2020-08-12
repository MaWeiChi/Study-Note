package main

import (
	"github.com/gin-gonic/gin"
)

type CheckaliveEntryList struct {
	List []CheckaliveEntry `json:"list,omitempty"`
}

type CheckaliveEntry struct {
	IntervalSec int    `json:"intervalSec,omitempty" validate:"required,min=10,max=86400"`
	Enable      bool   `json:"enable,omitempty"`
	TargetHost  string `json:"targetHost,omitempty" validate:"required,min=2,max=253,ipv4|fqdn"`
}

func main() {
	var chs CheckaliveEntryList
	var ch CheckaliveEntry
	ch.IntervalSec = 60
	chs.List = append(chs.List, ch)
	router := gin.Default()
	router.GET("", func(c *gin.Context) {

		c.JSON(200, gin.H{"data": chs, "count": len(chs.List)})
	})

	router.Run()
}
