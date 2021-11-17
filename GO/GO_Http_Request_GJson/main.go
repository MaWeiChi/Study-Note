package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

func main() {

	result, err := Request("GET", "http://10.123.12.144:58000/api/v1/dlm", "")
	if err != nil || gjson.GetBytes(result, "date.connectionStatus.status").String() != "connected" {
		fmt.Println(gjson.GetBytes(result, "data.connectionStatus.status").String())
		fmt.Println("OS patch", "OS patch", "OS patch schedule job skip", "Device", "Device", "OS patch schedule job skip. Reason: DLM Cloud disconnect.")
	}
	fmt.Println(gjson.GetBytes(result, "data.connectionStatus.status").String())
}

func Request(method string, url string, payload string) ([]byte, error) {
	client := &http.Client{}
	method = strings.ToUpper(method)
	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer client.CloseIdleConnections()
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}
