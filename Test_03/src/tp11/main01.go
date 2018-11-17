package main

import (
	"fmt"
)

type manager struct {
	name string
	age  int
}

func configInit() []*manager {

	managers := make([]*manager, 0, 10)

	managers = append(managers, &manager{name: "yzh", age: 18})
	managers = append(managers, &manager{name: "xjj", age: 16})
	managers = append(managers, &manager{name: "xyj", age: 20})
	managers = append(managers, &manager{name: "lcm", age: 22})

	return managers

}

func main() {

	for k, v := range configInit() {

		fmt.Printf("%+v %+v \n", k, v)

	}

	tmpp := *(configInit()[0])

	fmt.Printf("%T , %v %v %v\n", tmpp, tmpp, tmpp.name, tmpp.age)

}
