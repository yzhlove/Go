package main

import (
	"github.com/pkg/errors"
	"log"
)

// 初始化语句支持

func check(x int) error {
	if x <= 0 {
		return errors.New("x <= 0")
	}
	return nil
}


//func main() {
//
//	x := 10
//
//	if err := check(x) ; err == nil {
//		x++
//		println(x)
//	} else {
//		log.Fatalln(err)
//	}
//}

func main() {

	x := 10
	if err := check(x);err != nil {
		log.Fatalln(err)
	}
	x++
	println(x)
}



