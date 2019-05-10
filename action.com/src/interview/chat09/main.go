package main

import (
	"fmt"
	"sync"
)

type UserAgrs struct {
	peoples map[string]int
	sync.Mutex
}

func (ua *UserAgrs) Add(name string, age int) {
	ua.Lock()
	ua.peoples[name] = age
	ua.Unlock()
}

func (ua *UserAgrs) Get(name string) int {
	if age, ok := ua.peoples[name]; ok {
		return age
	}
	return -1
}

func main() {

	userAgrs := UserAgrs{
		peoples: make(map[string]int),
	}
	userAgrs.Add("yzh", 23)
	fmt.Println("Age = ", userAgrs.Get("yzh"))

}
