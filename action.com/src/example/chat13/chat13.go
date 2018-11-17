package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//从键盘读取输入

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("input Text:")

	input , err := inputReader.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("input :",input)

}
