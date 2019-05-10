package main

import (
	"encoding/json"
	"fmt"
)

// 通过匿名内部类解析json

type Screen struct {
	Size       float32
	ResX, ResY int
	Brand      string
}

type Battery struct {
	Capacity int
}

func generateJSON() []byte {

	jsonInfo := struct {
		Screen
		Battery
		HasTouch bool
	}{
		Screen: Screen{
			Size:  5.2,
			ResX:  3124,
			ResY:  2860,
			Brand: "Sony G8142",
		},

		Battery: Battery{
			Capacity: 3000,
		},

		HasTouch: true,
	}

	bytes, err := json.Marshal(jsonInfo)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func main() {

	getJson := generateJSON()
	fmt.Println("json_data:", string(getJson))

	screenAndTouch := struct {
		Screen
		HasTouch bool
	}{}

	_ = json.Unmarshal(getJson, &screenAndTouch)
	fmt.Println("screenAndTouch:", screenAndTouch)

	batteryAndTouch := struct {
		Battery
		HasTouch bool
	}{}

	_ = json.Unmarshal(getJson, &batteryAndTouch)
	fmt.Println("batteryAndTouch:", batteryAndTouch)

	allInfo := struct {
		Screen
		Battery
		HasTouch bool
	}{}

	_ = json.Unmarshal(getJson, &allInfo)
	fmt.Println("allInfo:", allInfo)

}
