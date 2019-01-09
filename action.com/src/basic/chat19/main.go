package main

import (
	"fmt"
	"math/rand"
	"time"
)

func eat() chan string {
	out := make(chan string)
	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		out <- "起床吃饭!"
		close(out)
	}()
	return out
}

func main() {

	eatch := eat()
	sleep := time.NewTimer(time.Second * 2)
	select {
	case s, ok := <-eatch:
		if ok {
			fmt.Println(s)
		}
	case <-sleep.C:
		fmt.Println("Time to sleep")
		//default:
		//	fmt.Println("Beat DouDou")
	}
}
