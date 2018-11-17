package main

import "fmt"

//数组

func main()  {

	a := [5]int{1,2,3,4,5}
	b := [...]int{1,2,3,4,5}
	c := [5]int{1:10,2:20}

	fmt.Printf("%v , %v ,%v \n",a,b,c)

}
