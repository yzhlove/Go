package main

import (
	"sync"
	"time"

	"fmt"
)

//go cond

var counter = 0

type event struct {
	mutex *sync.Mutex
	cond  *sync.Cond
}

func main() {

	et := new(event)
	et.mutex = new(sync.Mutex)
	et.cond = sync.NewCond(et.mutex)

	go wait(et)
	run(et)
	fmt.Println("Done .")
}

func run(et *event) {
	for i := 0; i < 100; i++ {
		func() {
			et.mutex.Lock()
			defer et.mutex.Unlock()
			counter++
			et.cond.Signal()
			time.Sleep(time.Second)
			fmt.Println("counter == ", counter)
		}()
	}
}

func wait(et *event) {
	for {
		func() {
			et.mutex.Lock()
			defer et.mutex.Unlock()
			for counter < 10 {
				et.cond.Wait()
				fmt.Println("event.")
			}
			counter = 0
			fmt.Println("ok.")
		}()
	}
}
