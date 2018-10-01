package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var data struct {
		Extra string
	}

	var data2 struct {
		T       string `json:"type"`
		Content string `json:"content"`
	}

	var data3 struct {
		SID string `json:"sid"`
		RID string `json:"game_history_id"`
	}

	data3.SID = "31sbaca89f9921a2a7777c1812d9f6e83f5b6"
	data3.RID = "7234"

	b, _ := json.Marshal(data3)

	data2.T = "login"
	data2.Content = string(b)

	b2, _ := json.Marshal(data2)

	data.Extra = string(b2)

	b3, _ := json.Marshal(data)

	fmt.Println(string(b3))
}
