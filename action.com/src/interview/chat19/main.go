package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover err :", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {
	defer p.deferError()
	for msg := range msgchan {
		m := msg.(int)
		fmt.Println("msg: ", m)
	}
}

func (p *Project) run(meaghan chan interface{}) {
	for {
		go p.exec(meaghan)
		time.Sleep(time.Second * 2)
	}
}

func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 10000)
}

func main() {
	p := new(Project)
	p.Main()
}
