package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {

	fmt.Println("BEFORE:", runtime.NumGoroutine())
	AllResp()
	fmt.Println("AFTER:", runtime.NumGoroutine())
	time.Sleep(time.Second)
}

func GetResp(i int) string {
	temp := []string{
		"xjj",
		"xyj",
		"lcm",
		"fyb",
	}
	return strconv.Itoa(i) + ":" + temp[rand.Intn(len(temp))]
}

func AllResp() {

	index := 10
	out := make(chan string)
	temp := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(index)
	go func() {
		//for tempValue := range out {
		//	fmt.Println(tempValue)
		//}
		for {
			select {
			case value, ok := <-out:
				if ok {
					fmt.Println(value)
				} else {
					fmt.Println("What Fuck.")
				}
			case <-temp:
				//close(out)
				goto EXIT
			}
		}
	EXIT:
	}()
	for i := 0; i < index; i++ {
		go func(i int) {
			defer wg.Done()
			out <- GetResp(i)
		}(i)
	}
	wg.Wait()
	temp <- struct{}{}
	//fmt.Println("Done .")
}
