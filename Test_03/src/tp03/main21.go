package main

import "log"

// 连续调用panic ，仅仅最后一个被捕获

func main() {

	defer func() {
		if err := recover();err != nil {
			log.Println(err)
		}  else {
			log.Fatalln(err)
		}
	}()
	defer func() {
		panic("i love you")
	}()
	panic("what are you going")

}

/*
2018/10/06 11:29:47 i love you
*/