package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/go-playground/validator"
)

func main() {
	var validate = validator.New()
	validate.RegisterValidation("netmask", validateNetmask)
	validate.RegisterValidation("enableDhcp", validateDHCP)
	validate.RegisterValidation("enableCheckalive", validateCheckalive)

	file, err := os.Open("/home/moxa/Study-Note/GO/GO_ValidatorRegister/config")
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

}

func validateNetmask(fl validator.FieldLevel) bool {
	netmaskRegexString := "^(?:(?:254|252|248|240|224|192|128).0.0.0|255.(?:254|252|248|240|224|192|128|0).0.0|255.255.(?:254|252|248|240|224|192|128|0).0|255.255.255.(?:254|252|248|240|224|192|128|0))$"
	netmaskRegex := regexp.MustCompile(netmaskRegexString)

	return netmaskRegex.MatchString(fl.Field().String())
}

func validateDHCP(fl validator.FieldLevel) bool {
	dhcp := fl.Field()
	ip := fl.Parent().FieldByName("Ip")
	netmask := fl.Parent().FieldByName("Netmask")
	if !dhcp.IsValid() || !ip.IsValid() || !netmask.IsValid() {
		return false
	}
	if dhcp.Bool() && (ip.String() == "" || netmask.String() == "") {
		return false
	}
	return true
}

func validateCheckalive(fl validator.FieldLevel) bool {
	checkalive := fl.Field()
	targetHost := fl.Parent().FieldByName("TargetHost")
	intervalSec := fl.Parent().FieldByName("IntervalSec")
	if !checkalive.IsValid() || !targetHost.IsValid() || !intervalSec.IsValid() {
		return false
	}
	if checkalive.Bool() && (intervalSec.IsZero() || targetHost.String() == "") {
		return false
	}
	return true
}

type CheckaliveEntry struct {
	IntervalSec        int    `json:"intervalSec" validate:"min=10,max=86400"`
	Enable             bool   `json:"enable" validate:"checkalive"`
	TargetHost         string `json:"targetHost" validate:"min=2,max=253,ipv4|fqdn"`
	Layer3Successfully bool   `json:"-"`
}

type NetworkLayerEntry struct {
	Name       string          `json:"name"`
	Checkalive CheckaliveEntry `json:"checkalive,omitempty" validate:"required"`
}
