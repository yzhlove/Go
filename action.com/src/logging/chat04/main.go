package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//json反序列化到变量

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

var JSON = `{
	"name":"Gopher",
	"title":"programmer",
	"contact":{
		"home":"415.333.333",
		"cell":"415.555.555"
	}
}`

//将JSON解析到结构体
func mainVar() {

	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Printf("contact = %v \n", c)
}

//将JSON解析到MAP变量
func main() {

	var c map[string]interface{}

	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Printf("Map:%T %v \n", c, c)
	fmt.Println("Name:", c["name"])
	fmt.Println("title", c["title"])
	fmt.Println("contact", c["contact"])
	fmt.Println("home", c["contact"].(map[string]interface{})["home"])
	fmt.Println("home", c["contact"].(map[string]interface{})["cell"])
}
