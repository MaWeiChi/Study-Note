package main

import (
	"fmt"
	"os"
)

func main() {
	os.MkdirAll("/home/moxa/Erik/Study-Note/GO/GO_Mkdir/1", 0755)
	_, err := os.Create("/home/moxa/Erik/Study-Note/GO/GO_Mkdir/1/log.db")
	if err != nil {
		fmt.Print(err.Error())
	}
	// os.MkdirAll(filepath.Dir("/home/moxa/Erik/Study-Note/GO/GO_Mkdir/1/log.db"), 0755)
}
