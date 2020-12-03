package main

import (
	"fmt"
	"io/ioutil"
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

	b, _ := ioutil.ReadFile("/home/moxa/Study-Note/GO/GO_IO&Read/operstate1")
	fmt.Println(strings.TrimSpace(string(b)))

	// b, _ := ioutil.ReadFile("/home/moxa/Study-Note/GO/GO_IO&Read/carrier")
	// fmt.Println(strings.TrimSpace(string(b)))
	// fmt.Println(b)
}
