package main

import (
	"encoding/json"
	"log"

	"github.com/tidwall/sjson"
)

const json1 = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
const json2 = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	var t Test
	// ts := []Test{}
	t.ID = 1
	t.Name = "abc"
	CovertToJSONString(t)
	// ts = append(ts, t)
	value, _ := sjson.Set("", "name.last", "Anderson")
	println(value)
	value, _ = sjson.Set("", "sub.0", t)

	println(value)
}

func CovertToJSONString(data interface{}) string {
	js, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error: %s", err)
		return ""
	}
	return string(js)
}

type Test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
