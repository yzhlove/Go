package main

import (
	"fmt"
	"sync"
	"time"
)

//锁尽量传递references，而不传递value

type data struct {
	sync.Mutex
}

//使用锁等资源的时候，传递指针
func (d *data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	var d data

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("wait")
	}()

	wg.Wait()

}
