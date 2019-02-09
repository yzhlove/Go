package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//select 的使用

func generate() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			i++
			out <- i
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		}
	}()
	return out
}

func createWorker(id int) chan<- int {
	work := make(chan int)
	go worker(id, work)
	return work
}

func worker(id int, c <-chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("id %d value:%d \n", id, n)
	}
}

func main() {

	var c1, c2 = generate(), generate()
	var work = createWorker(0)
	var values []int
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		var tp = 0
		var activeChannel chan<- int
		if len(values) > 0 {
			activeChannel = work
			tp = values[0]
		}
		select {
		case value := <-c1:
			values = append(values, value)
		case value := <-c2:
			values = append(values, value)
		case activeChannel <- tp:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue length = ", len(values))
		case <-after:
			fmt.Println("Bye !")
			os.Exit(0)
			return
		}
	}

}
