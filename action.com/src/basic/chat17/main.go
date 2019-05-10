package main

import (
	"fmt"
	"sync"
)

//流水线模型改良

//生产者
func productor(nums ...int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v
		}
	}()
	return out
}

func opt(t int) int {
	return t*2 + 1*3 + 4
}

//消费者

func calc(pt chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range pt {
			out <- opt(value)
		}
	}()
	return out
}

func merge(ch ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	tempFunc := func(tmp chan int) {
		defer wg.Done()
		for n := range tmp {
			out <- n
		}
	}
	wg.Add(len(ch))
	for _, c := range ch {
		go tempFunc(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	pt := productor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)
	c1 := calc(pt)
	c2 := calc(pt)
	c3 := calc(pt)
	c4 := calc(pt)

	for ret := range merge(c1, c2, c3, c4) {
		fmt.Println("ret = ", ret)
	}
	fmt.Println()
}
