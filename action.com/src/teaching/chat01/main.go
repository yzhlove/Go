package main

import "fmt"

func main() {

	var str = []string{"a", "b", "c", "d"}

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str); j++ {
			fmt.Println(str[i:i+1], str[0:j+1])
		}
	}

}
