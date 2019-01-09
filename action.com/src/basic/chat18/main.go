package main

import "fmt"

//select的使用

func main() {

	read := make(chan int, 1)
	write := make(chan int, 1)

	y := 1
	select {
	case x, ok := <-read:
		if ok {
			fmt.Println("x = ", x)
		}
	case write <- y:
		fmt.Println("write: ", y)
	default:
		fmt.Println("default ...")
	}

}
