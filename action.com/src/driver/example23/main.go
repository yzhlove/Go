package main

import (
	"fmt"
	"strings"
)

func main() {

	arrayList := []string{"User_Name:1", "User_Age:2", "User_Birthday:3", "User_Height:", "User_Width:", "User_Info"}

	for _, value := range arrayList {
		fmt.Println(strings.Split(value, ":"))
	}

}
