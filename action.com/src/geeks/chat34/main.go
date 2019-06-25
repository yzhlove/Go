package main

import "fmt"

//一个有趣的例子

func main() {

	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var counter1, counter2 int

	for i := 0; i < 1000; i++ {
		select {
		case <-c1:
			counter1++
		case <-c2:
			counter2++
		}
	}

	fmt.Println("Done counter1 = ", counter1, " counter2 = ", counter2)

}
