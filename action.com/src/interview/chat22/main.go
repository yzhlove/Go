package main

import "fmt"

type query func(string) string

func exec(name string, vs ...query) string {
	out := make(chan string)
	//fn := func(i int) {
	//	out <- vs[i](name)
	//}
	for i, _ := range vs {
		go func(i int) {
			str := vs[i](name)
			fmt.Println("str = ", str)
			out <- str
		}(i)
	}
	return <-out
}

func main() {
	ret := exec("111", func(s string) string {
		return s + "func1"
	}, func(s string) string {
		return s + "func2"
	}, func(s string) string {
		return s + "func3"
	}, func(s string) string {
		return s + "func4"
	})
	fmt.Println(ret)
}
