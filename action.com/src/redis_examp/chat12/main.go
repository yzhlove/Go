package main

import (
	"fmt"
	"sync"
)

func main() {

	if err := Init(); err != nil {
		panic(err)
	}

	fmt.Printf("========== over ==========")

}

func Init() error {

	var (
		wg     sync.WaitGroup
		errMsg = make(chan error, 5)
	)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				errMsg <- nil
			} else {
				errMsg <- fmt.Errorf("Err:%d ", i)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(errMsg)
	}()

	for msg := range errMsg {
		fmt.Printf("Msg = %v \n", msg)
		if msg != nil {
			return msg
		}
	}
	return nil
}
