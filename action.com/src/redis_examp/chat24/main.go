package main

import (
	"fmt"
	"strings"
)

func main() {

	list := []string{"yzh", "lcm", "xjj", "hxy", "fyb", "xyj"}
	Run("master", "xjj", list...)

}

func Run(name, key string, values ...string) {
	fmt.Println(name)
	fmt.Println(key)
	fmt.Printf("%T %v", values, values)

	if result := strings.Index(strings.Join(values, " "), key); result == -1 {
		fmt.Println("not found ")
	} else {
		fmt.Println("result = ", result)
	}

}
