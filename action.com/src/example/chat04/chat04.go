package main

import "fmt"

func main() {

	arrayString := [3]*string{new(string),new(string),new(string)}

	*arrayString[1] = "red"
	*arrayString[2] = "green"

	fmt.Printf("%v \n",arrayString)

	for index,value := range arrayString {
		if  *value == ""{
			fmt.Println("index : " , index)
		}
		fmt.Printf("%p %v \n",value,*value)
	}

}
