package main

import (
	"bufio"
	"fmt"
	"os"
)

//defer的使用

func tryDefer() {

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			// panic("exit...")		//panic -> defer之后出发
			fmt.Println("exit...")
			return
		}
	}

}

func tryDefer2() {

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	return

}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func writeFile(filename string) {

	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()

	f := fib()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(write, f())
	}

}

func main() {

	const filename = "/Users/love/WorkSpace/Go/action.com/src/demo/chat28/fib.txt"

	// tryDefer()

	// tryDefer2()

	writeFile(filename)

}
