package main

import (
	"fmt"
	"teaching/chat39/tool"
)

type Skill struct {
	Name  string
	Level int
}

type Actor struct {
	Name   string
	Age    int
	Skills []Skill
}

func main() {

	actor := Actor{
		Name: "yzh",
		Age:  18,
		Skills: []Skill{
			{Name: "Roll and roll", Level: 1},
			{Name: "Flash your dog eye", Level: 2},
			{Name: "Time to have Lunch", Level: 3},
		},
	}

	bytes, err := tool.ToJson(actor)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
