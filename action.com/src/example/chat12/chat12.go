package main

import "fmt"

//读取用户输入

func main() {

	var firstName,lastName string

	for {
		fmt.Scanln("%s %s",&firstName,&lastName)

		if firstName == lastName {
			break
		}
		fmt.Printf("%v %v \n",firstName,lastName)
	}

	fmt.Println("Done!")

}
