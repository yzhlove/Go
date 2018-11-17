package main

import (
	"exercise/worlds"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	filename := os.Args[1]
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)
	count := worlds.CountWorlds(text)

	fmt.Printf("worlds : %d \n", count)

}
