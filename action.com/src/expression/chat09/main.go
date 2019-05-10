package main

import (
	"fmt"
	"regexp"
)

//测试

var (
	gameRegex1 = regexp.MustCompile(`Game_Nickname+:[0-9]+`)
	userRegex1 = regexp.MustCompile(`User_[a-zA-Z]+:\{[0-9]+\}`)
	userRegex2 = regexp.MustCompile(`User_Platform:[a-zA-Z0-9]+:[0-9]+`)
)

func main() {

	keys := []string{
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
	}

	for _, key := range keys {

		switch {
		case userRegex1.MatchString(key):
			fmt.Println("User_1 ", key)
		case userRegex2.MatchString(key):
			fmt.Println("Platform ", key)
		case gameRegex1.MatchString(key):
			fmt.Println("Game ", key)
		default:
			fmt.Println("Default ", key)
		}
	}

}
