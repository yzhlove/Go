package main

import (
	"fmt"

	"github.com/davyxu/golog"
)

//测试

var log = golog.New("Project")

func tryRecover() {

	defer func() {

		err := recover()
		fmt.Printf("Err :%v \n", err)

		log.Debugf("Err:%v \n", err)

	}()

	panic("this is error")

}

func main() {

	tryRecover()

}
