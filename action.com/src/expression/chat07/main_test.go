package test

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

//场景 读速度大于写的速度

func Benchmark_ReadAndWrite(t *testing.B) {

	inChan := sourceChannel(t)
	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go workChannel(i, inChan, &wg)
	}
	wg.Wait()
	fmt.Println("Done .")
}

func sourceChannel(t *testing.B) chan string {
	inChan := make(chan string, 128)
	go func() {
		for i := 0; i < t.N; i++ {
			inChan <- "msg:" + strconv.Itoa(i)
			time.Sleep(50 * time.Millisecond)
		}
		close(inChan)
	}()
	return inChan
}

func workChannel(id int, inChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for range inChan {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("[%d] Done.\n", id)
}
