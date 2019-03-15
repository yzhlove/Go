package main

import (
	"testing"
)

func BenchmarkScheduler(b *testing.B) {

	index := b.N
	exit := make(chan bool)
	out := make(chan int)
	go func() {
		count := 0
		for {
			<-out
			count++
			if count == index {
				exit <- true
			}
		}
	}()
	for i := 0; i < index; i++ {
		go func() {
			out <- i
		}()
	}
	<-exit
}
