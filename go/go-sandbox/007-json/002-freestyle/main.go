package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const jsn = `{
	"users": [
		{
			"id": 1,
			"name": "Bob Great",
			"age": 123
		},
		{
			"id": 2,
			"name": "Mary Cool",
			"age": 34
		}
	]
}`

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var doc struct {
		Users []user `json:"users"`
	}

	if err := json.NewDecoder(strings.NewReader(jsn)).Decode(&doc); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", doc.Users)
	}
}
