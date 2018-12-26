package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//编码JSON

func main() {

	data := make(map[string]interface{})
	data["name"] = "yzh"
	data["title"] = "Programmer"
	data["contact"] = map[string]interface{}{
		"home":  "0724-7480611",
		"phone": "13617247924",
	}

	str, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(string(str))

}
