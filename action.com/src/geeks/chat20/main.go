package main

import (
	"sync"
	"time"

	"fmt"
)

//go cond

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := &Button{Clicked: sync.NewCond(&sync.Mutex{})}
	subscribe := func(c *sync.Cond, fn func()) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			wg.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		wg.Wait()
	}

	var clickRegister sync.WaitGroup
	clickRegister.Add(3)

	subscribe(button.Clicked, func() {
		defer clickRegister.Done()
		fmt.Println("- A -")
	})

	subscribe(button.Clicked, func() {
		defer clickRegister.Done()
		fmt.Println("- B -")
	})

	subscribe(button.Clicked, func() {
		defer clickRegister.Done()
		fmt.Println("- C -")
	})

	time.Sleep(time.Second)

	button.Clicked.Broadcast()

	clickRegister.Wait()

}
