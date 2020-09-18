package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator"
)

type RegisterReq struct {
	// gt = greater than
	Username    string `validate:"gt=3"`
	PasswordNew string `validate:"gt=3,password"`
	// eqfield = PasswordRepeat equal PasswordNew
	PasswordRepeat string `validate:"eqfield=PasswordNew"`
	// email format
	Email string `validate:"email"`
}

// type SecurityEntry struct {
// 	Mode       string `json:"mode" validate:"omitempty,eq=wpa2"`      // only wpa2 now
// 	Encryption string `json:"encryption" validate:"omitempty,eq=aes"` // only aes now
// 	Password   string `json:"password" validate:"omitempty,min=8,max=63,printascii"`
// 	Support    *bool  `json:"support,omitempty"`
// }

type BigSecurityEntry struct {
	Type     string        `json:"type"`
	Security SecurityEntry `json:"security"`
}

var validate = validator.New()

func validateFunc(req RegisterReq) error {
	err := validate.Struct(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("all clear")
	return err
}

func validatePassword(fl validator.FieldLevel) bool {
	if !strings.ContainsAny("0123456789", fl.Field().String()) || !strings.ContainsAny("~!@#$%^&*-_|;:,.<>[]{}()", fl.Field().String()) || len(fl.Field().String()) < 8 {
		return false
	}
	return true
}

func main() {
	/*
		validate.RegisterValidation("password", validatePassword)
		var req RegisterReq
		req.Username = "1"
		fmt.Println(req)
		// PasswordNew = PasswordRepeat = ""
		validateFunc(req)

		req.Username = "1234"
		fmt.Println(req)
		validateFunc(req)
		req.PasswordNew = "11234"
		req.Email = "example@example.com"
		fmt.Println(req)
		validateFunc(req)
		req.PasswordNew = "~11234#11"
		req.PasswordRepeat = "~11234#11"
		fmt.Println(req)
		validateFunc(req)

		if strings.ContainsAny(numbers, req.Username) {
			fmt.Println("A")
		}
		if strings.ContainsAny(special, req.Username) {
			fmt.Println("B")
		}
	*/

	/*
		file, err := os.Open("/home/moxa/Study-Note/GO/GO_Validator/config")
		if err != nil {
			panic(err)
		}
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		var entry SecurityEntry
		if err = json.Unmarshal(bytes, &entry); err == nil {
			err = validate.Struct(entry)
			fmt.Println(err)
		}
		var bigSecurity BigSecurityEntry
		bigSecurity.Security = entry
		err = validate.Struct(bigSecurity)
		fmt.Println(err)
	*/

	file, err := os.Open("/home/moxa/Study-Note/GO/GO_Validator/configwifi")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var entry WifiEntry
	if err = json.Unmarshal(bytes, &entry); err == nil {
		err = validate.Struct(entry)
		fmt.Println(err)
	}
	fmt.Println(entry.Client)
	err = validate.Struct(entry.Client.Networks)
	fmt.Println(err)
	err = validate.Struct(entry.Client.Networks[1])
	fmt.Println(err)
	// var t ClientSecurityEntry
	// t.Mode = entry.Client.Networks[1].Security.Mode
	// fmt.Println(t.Mode)
	// err = validate.Struct(t)
	// fmt.Println(err)
}

type WifiClientEntry struct {
	ConnectState string              `json:"connectState"`
	CurrentAp    string              `json:"currentAp"`
	Priority     []string            `json:"priority"`
	Networks     []ClientServerEntry `json:"networks" validate:"dive"`
}

// ClientServerEntry ...
type ClientServerEntry struct {
	Uuid      string              `json:"uuid"`
	Ssid      string              `json:"ssid"`
	Bssid     string              `json:"bssid"`
	LockBssid bool                `json:"-"`
	Security  ClientSecurityEntry `json:"security" `
	Signal    int16               `json:"signal"`
	Band      string              `json:"band"`
	wpaId     int
}

// type SecurityEntry struct {
// 	Mode       string `json:"mode" validate:"omitempty,eq=wpa2"`      // only wpa2 now
// 	Encryption string `json:"encryption" validate:"omitempty,eq=aes"` // only aes now
// 	Password   string `json:"password" validate:"omitempty,min=8,max=63,printascii"`
// 	Support    *bool  `json:"support,omitempty"`
// }

type ClientSecurityEntry struct {
	Mode       string `json:"mode" validate:"omitempty,oneof='wpa3' non"`
	Encryption string `json:"encryption" validate:"omitempty,oneof=aes tkip none"`
	Password   string `json:"password" validate:"omitempty,min=8,max=63,printascii"`
	Support    *bool  `json:"support,omitempty"`
}

// ScanListEntry ...
type ScanListEntry struct {
	Networks []ScanEntry `json:"networks"`
}

// ScanEntry ...
type ScanEntry struct {
	Ssid     string              `json:"ssid"`
	Security ClientSecurityEntry `json:"security"`
	Signal   int16               `json:"signal"`
	Band     string              `json:"band"`
}

type ClientEntry struct {
	Mac string `json:"mac"`
}

type SecurityEntry struct {
	Mode       string `json:"mode" validate:"omitempty,eq=wpa2"`      // only wpa2 now
	Encryption string `json:"encryption" validate:"omitempty,eq=aes"` // only aes now
	Password   string `json:"password" validate:"omitempty,min=8,max=63,printascii"`
	Support    *bool  `json:"support,omitempty"`
}

type ApEntry struct {
	Band          string        `json:"band" validate:"omitempty,oneof=band24 band50"`
	BroadcastSsid bool          `json:"broadcastSsid"`
	Ssid          string        `json:"ssid" validate:"omitempty,min=1,max=32,alphanum|containsany=-_"`
	Channel       int           `json:"channel"`
	Region        string        `json:"region"`
	Security      SecurityEntry `json:"security"`
	Clients       []ClientEntry `json:"clients"`
}

type WifiReadOnlyEntry struct {
	Type         string          `json:"type"`
	Name         string          `json:"name"`
	DisplayName  string          `json:"displayName"`
	Available    bool            `json:"available"`
	Status       bool            `json:"status"`
	Capabilities CapabilityEntry `json:"capabilities"`
	Regions      []RegionEntry   `json:"regions"`
	Interface    string          `json:"-"`
	Phy          string          `json:"-"`
	ctx          context.Context
	cancel       context.CancelFunc
}

type CapabilityEntry struct {
	Mode []string `json:"mode"`
	Band []string `json:"band"`
}

type RegionEntry struct {
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
	Band24      []int  `json:"band24"`
	Band50      []int  `json:"band50"`
}

type WifiReadWriteEntry struct {
	Id     int             `json:"id,omitempty" validate:"omitempty,min=1"`
	Enable bool            `json:"enable"`
	Mode   string          `json:"mode,omitempty" validate:"omitempty,oneof=ap client"` // only ap mode now
	Ap     ApEntry         `json:"ap,omitempty"`
	Client WifiClientEntry `json:"client,omitempty" validate:"dive"`
}

type WifiEntry struct {
	WifiReadOnlyEntry
	WifiReadWriteEntry
	logger *log.Logger
}
