package main

import "fmt"

func main() {

	s := make([]int,0,100)
	s1 := s[:2:4]
	s2 := append(s1,12,3,4,5,6,7)

	fmt.Printf("s1: %p:%v \n",&s1[0],s1)
	fmt.Printf("s1: %p:%v \n",&s2[0],s2)

	fmt.Printf("s data: %v\n",s[:10])
	fmt.Printf("s1 cap:%d ,s2 cap:%d \n ",cap(s1),cap(s2))
	
}
