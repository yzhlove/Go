package main

import (
	"fmt"
	"time"
)

//携程的使用

func task(id int) {
	fmt.Printf("%+v \n", id)
}

func main() {

	for i := 1; i <= 10; i++ {
		go task(i)
	}

	time.Sleep(2 * time.Second)

}
