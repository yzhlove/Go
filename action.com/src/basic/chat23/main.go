package main

import "fmt"

//Go select

//求值顺序 自上而下 从左到右

var ch1 chan int
var ch2 chan int
var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]
}

func main() {

	select {
	case getChan(0) <- getNumber(2):
		fmt.Println("One Read ...")
	case getChan(1) <- getNumber(3):
		fmt.Println("Two Read ...")
	default:
		fmt.Println("default")
	}
}
