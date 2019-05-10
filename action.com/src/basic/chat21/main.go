package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁的应用

type Bank struct {
	mutex  sync.RWMutex
	saving map[string]int
}

func NewBank() *Bank {
	b := &Bank{
		saving: make(map[string]int),
	}
	return b
}

func (b *Bank) Deposit(name string, amout int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if _, ok := b.saving[name]; !ok {
		b.saving[name] = 0
	} else {
		b.saving[name] += amout
	}
}

func (b *Bank) Withdraw(name string, amount int) int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if _, ok := b.saving[name]; !ok {
		return 0
	}
	if b.saving[name] < amount {
		amount = b.saving[name]
	}
	b.saving[name] -= amount
	return amount
}

//func (b *Bank) Query(name string) int {
//	b.mutex.Lock()
//	if _, ok := b.saving[name]; !ok {
//		return 0
//	}
//	b.mutex.Unlock()
//	return b.saving[name]
//}

func (b *Bank) Query(name string) int {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	if _, ok := b.saving[name]; !ok {
		return 0
	}
	return b.saving[name]
}

func main() {
	b := NewBank()
	go b.Deposit("xxx", 100)
	go b.Withdraw("xxx", 20)
	go b.Deposit("xxx", 2000)

	time.Sleep(2 * time.Second)
	fmt.Printf("has :%d \n", b.Query("xxx"))
	fmt.Printf("has :%d \n", b.Query("xxx"))
	fmt.Printf("has :%d \n", b.Query("xxx"))

}
