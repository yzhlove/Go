package main

import (
	"fmt"
	"runtime"
	"sync"
)

//临界资源区
//线程竞态

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)
	go inCounter(1)
	go inCounter(2)
	wg.Wait()

	fmt.Println("exitiing ... counter:", counter)

}

func inCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := counter
		//当前goroutine从线程退出,并放回队列
		runtime.Gosched()
		value++
		counter = value
	}
}
