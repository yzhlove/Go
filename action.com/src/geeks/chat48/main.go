package main

import "fmt"

//pipeline

func main() {

	repeat := func(done <-chan struct{}, values ...interface{}) <-chan interface{} {
		stream := make(chan interface{})
		go func() {
			defer close(stream)
			for _, it := range values {
				select {
				case stream <- it:
				case <-done:
					return
				}
			}
		}()
		return stream
	}

	take := func(done <-chan struct{}, ins <-chan interface{}, num int) <-chan interface{} {
		stream := make(chan interface{})
		go func() {
			defer close(stream)
			for i := 0; i < num; i++ {
				select {
				case stream <- ins:
				case <-done:
					return
				}
			}
		}()
		return stream
	}

	done := make(chan struct{})
	for num := range take(done, repeat(done, 1), 5) {
		fmt.Printf("num =%T %v \n", num, num)
	}
	close(done)
	fmt.Println("Done .")
}
