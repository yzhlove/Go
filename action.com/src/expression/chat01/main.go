package main

import (
	"fmt"
	"regexp"
)

//正则表达式测试
var keys = `
	"User_Platform:Default:24",
	"User_Employee:{1118434724826583040}:2008",
	"User_Equipment:{1118436117167738880}:1118436118321172480",
	"User_Quests:{1118434725627695104}",
	"Game_Nickname:0",
	"User_Account:{1118434726672076800}",
	"User_Res:{1118434726844043264}",
	"User_Equipments:{1118434727984893952}",
	"User_Furnitures:{1118434726311366656}",
	"User_Employees:{1118434726227480576}",
`

func main() {

	mustCompile := "(([a-zA-Z]+_[a-zA-z]+):{([0-9]+)}[:]?[0-9]*)"
	reg := regexp.MustCompile(mustCompile)

	result := reg.FindAllSubmatch([]byte(keys), -1)
	for _, matches := range result {
		for _, v := range matches {
			fmt.Println(string(v))
		}
		fmt.Println()
	}

}
