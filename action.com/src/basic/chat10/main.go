package main

import "fmt"

//slice操作
func main() {

	seq := []string{"a", "b", "c", "d", "e"}

	index := 2
	fmt.Println(seq[:index], seq[index+1:])
	fmt.Printf("%v \n", seq)
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Printf("%v \n", seq)

}
