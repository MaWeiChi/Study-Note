package main

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/vishvananda/netlink"
)

func getEthName() []string {
	reEnoN := regexp.MustCompile(`eno[0-9]`)
	reEthN := regexp.MustCompile(`eth[0-9]`)
	ethList := []string{}

	out, _ := exec.Command("cat", "/home/moxa/Study-Note/GO/GO_Netlink/interfaces").Output()
	for _, l := range strings.Split(string(out), "\n") {
		line := strings.TrimSpace(l)
		splits := strings.Split(line, " ")

		if splits[0] != "iface" {
			continue
		}
		if reEnoN.MatchString(splits[1]) || reEthN.MatchString(splits[1]) {
			ethList = append(ethList, splits[1])
		}
	}
	return ethList
}

func linkList() []netlink.Link {
	links, _ := netlink.LinkList()
	return links
}

func linkByName(name string) netlink.Link {
	link, _ := netlink.LinkByName(name)
	return link
}
