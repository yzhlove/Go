package main

import (
	"errors"
	"fmt"
)

//  函数返回值

func div(a,b int) (int ,error) {
	if b == 0 {
		return 0,errors.New("除数不能为0")
	}
	return a / b ,nil
}

func main()  {
	a,b := 1,2
	e,g := 2,0
	c,err := div(a,b)
	fmt.Println(c,err)

	c,err = div(e,g)
	fmt.Println(c,err)
}
