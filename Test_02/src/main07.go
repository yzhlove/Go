package main

import "fmt"

// 结构体

type user struct {
	name string
	age int
}

type manager struct {
	user
	title string
}

func main()  {
	var m manager
	m.name = "Tom"
	m.age = 22
	m.title = "CTO"
	fmt.Println(m)
}