package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//信号处理

var exits = &struct {
	sync.RWMutex
	funcs   []func()
	signals chan os.Signal
}{}

func atexit(f func()) {
	exits.Lock()
	exits.Unlock()
	exits.funcs = append(exits.funcs, f)
}

func waitExist() {
	if exits.signals == nil {
		exits.signals = make(chan os.Signal)
		signal.Notify(exits.signals, syscall.SIGINT, syscall.SIGTERM)
	}

	exits.RLock()
	for _, f := range exits.funcs {
		defer f()
	}
	exits.RUnlock()

	<-exits.signals

}

func main() {

	atexit(func() { fmt.Println("exit_1 ... ") })
	atexit(func() { fmt.Println("exit_2 ... ") })

	waitExist()

}
