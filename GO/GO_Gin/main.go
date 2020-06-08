package main

import (
	"github.com/gin-gonic/gin"
)

type Array struct {
	Array []string
}

func main() {
	var arr Array
	arr.Array = []string{}
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {

		c.JSON(200, gin.H{"data": arr})
	})

	router.Run()
}
