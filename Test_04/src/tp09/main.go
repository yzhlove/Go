package main

import (
	"fmt"
	"sync"
	"time"
)

//进程推出方式 信号量

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) //写在外面
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Second)

			fmt.Println("go run id = ", id)
		}(i)

	}

	fmt.Println("master run before ...")
	wg.Wait()
	fmt.Println("master run after ...")

}
