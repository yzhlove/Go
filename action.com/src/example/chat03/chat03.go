package main

import "fmt"

//数组复制



func main() {

	array1 := [5]string{}

	array2 := [5]string{"red","origin","yellow","green","blue"}

	fmt.Printf("%v \n",array2)

	array1 = array2

	fmt.Printf("%v \n",array1)

}
