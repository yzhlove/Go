package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type procInfo struct {
	id   int
	name string
}

func getStr() string {
	str := ""
	for i := 1; i <= 5; i++ {
		str += string(rune(rand.Intn(26) + 65))
	}
	return str
}

func main() {

	var wg sync.WaitGroup
	var mutex sync.Mutex
	mmp := make(map[int]*procInfo)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			mmp[id] = &procInfo{
				id:   id,
				name: getStr(),
			}
			fmt.Printf("%+v \n", mmp[id])
		}(i)
	}

	wg.Wait()
	fmt.Println("==================================")
	fmt.Printf("%+v \n", mmp)

}
