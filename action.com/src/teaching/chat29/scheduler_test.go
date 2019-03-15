package main

import "testing"

func BenchmarkSchedule(b *testing.B) {
	const N = 10
	out := make(chan chan int, N)
	exit := make(chan bool)
	count := b.N
	go func() {
		index := 0
		for {
			tmp := <-out
			tmp <- 1
			index++
			if index == count {
				exit <- true
			}
		}
	}()
	for i := 0; i < count; i++ {
		ch := make(chan int)
		out <- ch
		go func() {
			<-ch
		}()
	}
	<-exit
}
