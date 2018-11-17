package main

import (
	"time"
)

//Go中的并发操作

func main() {

	m := make(map[string]int)

	go func() {
		for {
			m["a"] += 1
		}
	}()

	go func() {
		for {
			_ = m["b"]
			time.Sleep(time.Microsecond)
		}
	}()

	select {}

}
