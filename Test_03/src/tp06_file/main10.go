package main

import (
	"fmt"
)

func main() {

	type user struct {
		Name string
		Age  int
	}

	user1 := user{Name: "Tom", Age: 20}
	user2 := user{Name: "Microsoft", Age: 21}

	fmt.Println("user1 = ", user1, " user2 = ", user2)

}
