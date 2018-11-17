package main

import "fmt"

// rnue string byte之间的准换

func main() {

	r := '我'

	s := string(r)
	b := byte(r)

	s2 := string(b)
	r2 := rune(b)

	fmt.Println(s,b,s2,r2)

}
