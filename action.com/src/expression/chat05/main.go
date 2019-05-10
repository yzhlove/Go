package main

import "fmt"

func main() {

	switch {
	case Test(1):
		fmt.Println(1)
	case Test(2):
		fmt.Println(2)
	case Test(3):
		fmt.Println(3)
	}

	fmt.Println("ok")

}

func Test(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}
