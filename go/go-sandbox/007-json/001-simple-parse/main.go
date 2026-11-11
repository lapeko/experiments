package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

func main() {
	jsn := `{
		"users": [
			{"id": 1, "name": "Alice", "active": null},
			{"id": 2, "name": "Bob", "active": false}
		]
	}`

	var users struct {
		Users []user `json:"users"`
	}
	if err := json.Unmarshal([]byte(jsn), &users); err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Printf("%v", users)
	}
}
