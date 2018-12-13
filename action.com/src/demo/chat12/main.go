package main

import (
	"fmt"
	"runtime"
	"sync"
)

//一个简单的并发

func main() {

	//限制最多在一个逻辑处理器上运行
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}

	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Printf("exitting...")

}
