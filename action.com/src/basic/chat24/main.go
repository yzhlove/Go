package main

import "fmt"

//如果右多个case满足条件，则随机执行一个

func main() {

	tempIndex := 5
	cp := make(chan string, tempIndex)

	for i := 0; i < tempIndex; i++ {
		select {
		case cp <- fmt.Sprintf("[%d] Send Message", 1):
		case cp <- fmt.Sprintf("[%d] Send Message", 2):
		case cp <- fmt.Sprintf("[%d] Send Message", 3):
		}
	}

	for i := 0; i < tempIndex; i++ {
		fmt.Println("Read Message: ", <-cp)
	}

}
