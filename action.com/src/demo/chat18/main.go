package main

import (
	"fmt"
	"runtime"
	"sync"
)

//修改临街资源区

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex //申明并初始化一个互斥锁
)

func main() {

	wg.Add(2)

	go inCount("A")
	go inCount("B")

	wg.Wait()

	fmt.Printf("counter = %v \n", counter)
	fmt.Printf("exiting ... \n")
}

func inCount(name string) {

	defer wg.Done()

	for count := 0; count < 2; count++ {
		//临界资源区，加锁，保证协程执行的正确性
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
	fmt.Printf("%s is Down!", name)

}
