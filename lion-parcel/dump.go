package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	payload := struct {
		Id     string
		Name   string
		Detail struct {
			ItemID string
			LogID  string
		}
	}{

		Id:   "111",
		Name: "111",
		Detail: struct {
			ItemID string
			LogID  string
		}{
			ItemID: "111",
			LogID:  "111",
		},
	}
	fmt.Println("before", payload)
	fmt.Println("after", DumpToString(payload))
}

func DumpToString(v interface{}) string {

	str, ok := v.(string)
	if !ok {
		buff := &bytes.Buffer{}
		json.NewEncoder(buff).Encode(v)
		return buff.String()
	}

	return str
}
