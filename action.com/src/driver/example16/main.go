package main

import "fmt"

func main() {

	arrays := []string{}

	temp := []string{"a", "b", "c"}

	arrays = append(arrays, temp...)
	fmt.Println(arrays)
}
