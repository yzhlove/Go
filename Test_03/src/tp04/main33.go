package main

import "fmt"

// map的使用

func main() {

	m := map[string]int{
		"a":1,
		"b":2,
	}

	fmt.Printf("%+v \n",m)

	m["a"] = 10
	m["b"] = 20

	fmt.Printf("%+v \n",m)

	if v,ok  := m["d"];ok {
		println("v = ",v)
	}

	delete(m,"d")

}
