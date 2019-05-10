package main

import "fmt"

type Profile struct {
	Name    string
	Age     int
	Married bool
}

var mapper = make(map[int][]*Profile)

func simpleHash(str string) int {
	ret := 0
	for i := 0; i < len(str); i++ {
		ret += int(str[i])
	}
	return ret
}

type searchKey struct {
	Name string
	Age  int
}

func (s *searchKey) getHash() int {
	return simpleHash(s.Name) + s.Age*1000000
}

func buildIndex(list []*Profile) {
	for _, pf := range list {
		key := searchKey{Name: pf.Name, Age: pf.Age}
		resultValue := mapper[key.getHash()]
		resultValue = append(resultValue, pf)
		mapper[key.getHash()] = resultValue
	}
}

func queryData(name string, age int) {
	key := searchKey{Name: name, Age: age}
	resultList := mapper[key.getHash()]
	for _, pf := range resultList {
		if pf.Name == name && pf.Age == age {
			fmt.Printf("Find :%v", pf)
			return
		}
	}
	fmt.Println("not found")
}

func main() {
	list := []*Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "里斯", Age: 25, Married: true},
		{Name: "王五", Age: 21},
	}
	buildIndex(list)
	queryData("张三", 30)

}
