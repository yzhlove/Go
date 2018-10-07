package main

import (
	"fmt"
	"github.com/pkg/errors"
)

//  栈

func main() {

	stack := make([]int, 0, 5)

	// 入栈
	push := func(x int) error {
		n := len(stack)
		if n == cap(stack) {
			return errors.New("stack is full")
		}
		stack = stack[:n+1]
		stack[n] = x
		return nil
	}

	// 出栈
	pop := func() (int ,error) {
		n := len(stack)
		if n == 0 {
			return 0,errors.New("stack is empty")
		}
		x := stack[n-1]
		stack = stack[:n-1]
		return x ,nil
	}


	for i := 0;i < 10;i++ {
		fmt.Printf("push %d:%v ,%v \n",i,push(i),stack)
	}

	for i := 0; i < 10;i++ {
		x,err := pop()
		fmt.Printf("pop %d, %v ,%v \n",x,err,stack)
	}

}
