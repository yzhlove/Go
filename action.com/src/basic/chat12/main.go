package main

import "fmt"

//匿名函数与回掉

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func main() {
	visit([]int{1, 2, 3, 4, 5}, func(value int) {
		fmt.Println("value = ", value)
	})
}
