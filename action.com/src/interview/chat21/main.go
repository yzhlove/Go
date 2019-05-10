package main

import "fmt"

type Student struct {
	name string
}

func main() {
	//m := map[string]Student{"people", {"zhoujielun"}}
	//m["people"].name = "wuyanzu"

	//m := map[string]*Student{"people": new(Student)}
	//m["people"].name = "zhoujielun"
	//fmt.Println(m)

	m := map[string]*Student{"people": &Student{name: "zhoujielun"}}
	m["people"].name = "wuyanzu"
	fmt.Println(m["people"])

}
