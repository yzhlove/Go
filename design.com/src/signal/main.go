package main

import (
	"sync"
)

//单利模式

type singletion map[string]string

var (
	once     sync.Once
	instance singletion
)

func main() {

}
