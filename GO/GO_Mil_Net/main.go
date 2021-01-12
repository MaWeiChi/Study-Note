package main

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {

	ipt := NewIPTools(nil, log.New())
	lines := strings.Split(ipt.ShowRoute("0.0.0.0/0"), "\n")
	fmt.Println(lines)
	for _, line := range lines {
		splits := strings.Split(line, " ")
		if len(splits) < 5 {
			continue
		}
		if splits[0] == "default" {
			fmt.Println(splits[2])
			fmt.Println(splits[4])

		}

	}
}
