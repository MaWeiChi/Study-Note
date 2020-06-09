package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Array struct {
	Array []string
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Person      *Person   `json:"person"`
	CreatedDate time.Time `json:"created_date"`
}

func main() {
	var eee SystemConfigurationBody
	eee.FileName = "123"
	eee.FileSize = 123
	eee.Diff = &DiffEntry{}
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {

		c.JSON(200, gin.H{"data": eee})
	})

	router.Run()
}

type SystemConfigurationBody struct {
	FileName string     `json:"fileName"`
	FileSize int64      `json:"fileSize"`
	Diff     *DiffEntry `json:"diff"`
}

type DiffEntry struct {
	Imported *ModelAndTPEVersionInfo `json:"imported",omitempty`
	Current  *ModelAndTPEVersionInfo `json:"current",omitempty`
	St       *string                 `json:"st"`
}

type ModelAndTPEVersionInfo struct {
	Model      string `json:"model",omitempty`
	TPEVersion string `json:"tpeversion",omitempty`
}
