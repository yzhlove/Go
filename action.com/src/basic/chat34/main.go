package main

import "fmt"

func generatePlayer(name string) func() (string, int) {
	hp := 150
	return func() (s string, i int) {
		s = name
		i = hp
		return
	}
}

func main() {

	generate := generatePlayer("yzh")
	name, hp := generate()
	fmt.Printf("name: %s hp:%d \n", name, hp)

	//fmt测试
	fmt.Println(5, "Hello", &struct {
		v int
		t string
	}{
		v: 123,
		t: "love",
	}, true)

}
