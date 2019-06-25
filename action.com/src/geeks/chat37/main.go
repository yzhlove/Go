package main

import (
	"bytes"
	"fmt"
	"sync"
)

//传入切片

func main() {

	show := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buf bytes.Buffer
		for _, b := range data {
			_, _ = fmt.Fprintf(&buf, "%c", b)
		}
		fmt.Println("value = ", buf.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go show(&wg, data[:3])
	go show(&wg, data[3:])
	wg.Wait()

	fmt.Println("Done .")
}
