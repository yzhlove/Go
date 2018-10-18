package main

//  go json_encode

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {

	Name string
	Age int
	Roles []string
	Skill map[string]float64
}


func main() {

	skill := make(map[string]float64)

	skill["python"] = 99.5
	skill["java"] = 90
	skill["ruby"] = 80.0

	user := User{
		Name:"rsj217",
		Age:27,
		Roles:[]string{"Owner","Master"},
		Skill:skill,
	}

	rs,err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(rs))


	var tempUser User
	err = json.Unmarshal(rs, &tempUser)
	if err != nil {
		log.Fatalln("解析失败!")
	} else {
		fmt.Println("User:", tempUser)
	}


}