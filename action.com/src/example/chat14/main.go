package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("input Text:")
		input , err := inputReader.ReadString('\n')
		input = strings.Replace(input,"\n","",-1)
		if input == "quit" ||  err != nil {
			break
		}
		fmt.Printf("[%v]\n",input)
	}
	fmt.Println("Done!")


}
