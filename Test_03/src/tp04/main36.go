package main

import "fmt"

// map的初始化

func main() {

	var m map[string]int		// nil
	m2 := map[string]int{}		// ""

	println(m["a"])	// ok
	//m["a"] = 10		// ERROR	panic: assignment to entry in nil map

	m2["a"] = 20	// ok
	fmt.Printf("%+v \n",m2)
	println(m == nil,m2 == nil)

}
