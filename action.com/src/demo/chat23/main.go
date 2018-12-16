package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//go 协程调度猜想

var (
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func main() {

	runtime.GOMAXPROCS(4)

	//使用Goruntine尽量避免无锁设计

	wg.Add(5)
	go task("A", 10)
	go task("B", 1)
	go task("C", 20)
	go task("D", 5)
	go task("E", 8)
	wg.Wait()
	fmt.Printf("All Task Over Exiting...")
}

func task(name string, ts int) {
	defer wg.Done()
	// mutex.Lock()
	fmt.Printf("%s start#%d time \n", name, ts)
	time.Sleep(time.Duration(ts) * time.Second)
	fmt.Printf("%s over#%d time \n", name, ts)
	// mutex.Unlock()
}
