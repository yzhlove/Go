package main

import "fmt"

// string -> byte

func main() {

	b := make([]byte,7)
	n := copy(b,"abcde")
	fmt.Println(n,b)

}
