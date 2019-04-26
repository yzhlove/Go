package main

import "fmt"

func main() {

	var distance = []string{"1", "2", "3", "4", "5"}

	source := make([]string, len(distance))

	copy(source, distance)

	fmt.Println(source)

}
