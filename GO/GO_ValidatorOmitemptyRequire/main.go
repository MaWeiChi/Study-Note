package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-playground/validator"
)

func main() {
	var validate = validator.New()

	file, err := os.Open("/home/moxa/Study-Note/GO/GO_ValidatorOmitemptyRequire/config")
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var entry NetworkLayerEntry
	if err = json.Unmarshal(bytes, &entry); err == nil {
		err = validate.Struct(entry)
		fmt.Println(err)
	}

	fmt.Println(entry.Name)
	fmt.Println(err)

	files, err := os.Open("/home/moxa/Study-Note/GO/GO_ValidatorOmitemptyRequire/configW")
	if err != nil {
		panic(err)
	}

	bytess, err := ioutil.ReadAll(files)
	if err != nil {
		panic(err)
	}

	var entrys NetworkLayerEntry
	if err = json.Unmarshal(bytess, &entrys); err == nil {
		err = validate.Struct(entrys)
		fmt.Println(err)
	}

	fmt.Println(entrys.Name)
	fmt.Println(err)
}

type CheckaliveEntry struct {
	IntervalSec        int    `json:"intervalSec" validate:"required,min=10,max=86400"`
	Enable             bool   `json:"enable"`
	TargetHost         string `json:"targetHost" validate:"required,min=2,max=253,ipv4|fqdn"`
	Layer3Successfully bool   `json:"-"`
}

type NetworkLayerEntry struct {
	Name       string          `json:"name"`
	Checkalive CheckaliveEntry `json:"checkalive,omitempty" validate:"omitempty"`
}
