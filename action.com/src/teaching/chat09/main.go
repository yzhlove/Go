package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	done = make(chan bool)
)

type Message struct {
	ID  int
	Msg string
}

func push(id int, head string, input chan Message) {
	count := 0
	for {
		count++
		msg := Message{
			ID:  id,
			Msg: head + strconv.Itoa(count),
		}
		input <- msg
		time.Sleep(time.Millisecond * 500)
	}
}

func sub() chan Message {
	out := make(chan Message)
	go func() {
		index := 0
		for {
			msg := <-out
			index++
			if index >= 100 {
				fmt.Printf("[last] [%d] message:%v \n", index, msg)
				done <- true
			}
			fmt.Printf("[read] [%d] message:%v \n", index, msg)
		}
	}()
	return out
}

func main() {

	in := sub()
	go push(1, "love_xyj:", in)
	go push(2, "love_lcm:", in)

	<-done
}
