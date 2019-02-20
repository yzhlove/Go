package main

import "fmt"

// 空接口类型比较

/*
map
slice
不能参与比较。会直接panic
*/

func main() {

	var a interface{} = 10
	var b interface{} = 20
	var c interface{} = 10

	fmt.Printf("%v %v \n", a == b, a == c)

}
