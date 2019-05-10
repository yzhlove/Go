package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Age  int
	Name string
}

type Temp struct {
	users []User
}

func main() {

	users := []User{
		{Age: 10, Name: "xjj"},
		{Age: 20, Name: "yzh"},
	}

	userBytes, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	//userBytes = []byte{}

	temp := &Temp{}

	if err := json.Unmarshal(userBytes, &temp.users); err != nil {
		panic(err)
	}

	fmt.Println(temp)

	var a []int
	fmt.Printf("%T %v %v \n", a, a, len(a))
	a = nil
	fmt.Printf("%T %v %v \n", a, a, len(a))
}
