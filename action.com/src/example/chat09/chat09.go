package main

import "fmt"

//切片追加

func main() {

	s1 := []int{1,2}
	s2 := []int{3,4}

	s2 = append(s2,s1...)

	fmt.Println(s2)

}
