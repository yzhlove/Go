package main

import (
	"fmt"
)

//Go Select

func main() {

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)

	select {
	case e1 := <-ch1:
		fmt.Printf("E1 = %d \n", e1)
	case e2 := <-ch2:
		fmt.Printf("E2 = %d \n", e2)
	default:
		fmt.Println("default !")
	}

}
