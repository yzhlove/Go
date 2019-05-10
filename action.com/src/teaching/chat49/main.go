package main

import "fmt"

type Profile struct {
	Name    string
	Age     int
	Married bool
}

type searchKey struct {
	Name string
	Age  int
}

//多键索引

var mapper = make(map[interface{}]*Profile)

func buildIndex(list []*Profile) {
	for _, pf := range list {
		key := searchKey{
			Name: pf.Name,
			Age:  pf.Age,
		}
		mapper[key] = pf
	}
}

func queryData(name string, age int) {
	key := searchKey{Name: name, Age: age}
	result, ok := mapper[key]
	if ok {
		fmt.Printf("Find :%v \n", result)
	} else {
		fmt.Println("not find ")
	}
}

func main() {
	list := []*Profile{
		{Name: "张三", Age: 33, Married: true},
		{Name: "里斯", Age: 25, Married: true},
		{Name: "王五", Age: 21},
	}
	buildIndex(list)
	queryData("张三", 33)
}
