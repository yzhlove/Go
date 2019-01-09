package main

import "fmt"

//流水线模型

func productor(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < len(nums); i++ {
			out <- nums[i]
		}
	}()
	return out
}

func square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()
	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := square(productor(nums...))
	for ret := range ch {
		fmt.Println("square value:= ", ret)
	}
	fmt.Println()
}
