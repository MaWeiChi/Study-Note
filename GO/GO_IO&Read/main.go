package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	// f, _ := os.Open("/home/moxa/Study-Note/GO/GO_IO&Read/operstate1")
	// defer f.Close()
	// f1, _ := os.Open("/home/moxa/Study-Note/GO/GO_IO&Read/operstate1")
	// defer f1.Close()

	// b := make([]byte, 10)
	// l, _ := f.Read(b)
	// for range b{{}
	// a := make([]byte, l)
	// f1.Read(a)
	// s := strings.TrimSpace(string(a[0:4]))
	// fmt.Println(s)

	// b, _ := ioutil.ReadFile("/home/moxa/Study-Note/GO/GO_IO&Read/operstate1")
	// fmt.Println(strings.TrimSpace(string(b)))

	// b, _ := ioutil.ReadFile("/home/moxa/Study-Note/GO/GO_IO&Read/carrier")
	// fmt.Println(strings.TrimSpace(string(b)))
	// fmt.Println(b)
	reEnpNs0 := regexp.MustCompile(`enp[0-9]s0`)
	reEthN := regexp.MustCompile(`eth[0-9]`)

	out, _ := exec.Command("cat", "/home/moxa/Study-Note/GO/GO_IO&Read/interfaces").Output()
	for _, l := range strings.Split(string(out), "\n") {
		line := strings.TrimSpace(l)
		splits := strings.Split(line, " ")

		if splits[0] != "iface" {
			continue
		}
		if reEnpNs0.MatchString(splits[1]) || reEthN.MatchString(splits[1]) {
			fmt.Println(splits[1])
		}

		// if strings.Contains(splits[0], "model name") {
		// 	dev.CPU = strings.TrimSpace(splits[1])
		// 	return
		// }
	}

	// out, _ := exec.Command("ls", "/sys/class/net/").Output()
	// for _, l := range strings.Split(string(out), "\n") {
	// 	// line := strings.TrimSpace(l)
	// 	// splits := strings.Split(line, " ")

	// 	// if splits[0] != "iface" {
	// 	// 	continue
	// 	// }
	// 	fmt.Println(l)
	// 	// if strings.Contains(splits[0], "model name") {
	// 	// 	dev.CPU = strings.TrimSpace(splits[1])
	// 	// 	return
	// 	// }
	// }
}
