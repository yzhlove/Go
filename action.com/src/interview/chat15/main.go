package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name"`
}

func main() {

	js := `{"name":"11"}`
	var p People
	if err := json.Unmarshal([]byte(js), &p); err != nil {
		panic(err)
	}
	fmt.Printf("%v", p)

	var people People
	people.Name = "11"

	if bytes, err := json.Marshal(people); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s", bytes)
	}

}
