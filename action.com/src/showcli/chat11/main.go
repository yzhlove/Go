package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	path := "./test.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println(file.Name())
	count := 0
	for {
		_, err := file.Seek(0, 1)
		if err != nil || err == io.EOF {
			break
		}

		//fmt.Println("cur = ", cur)
		buf := make([]byte, 1)
		_, err = io.ReadFull(file, buf)
		if err != nil || err == io.EOF {
			fmt.Println()
			fmt.Println("ERROR")
			break
		}
		if count%24 == 0 {
			if count == 0 {
				fmt.Print("[")
			} else {
				fmt.Printf("]\n[")
			}
		}

		if string(buf) == "\r" {
			fmt.Printf("*")
		} else if string(buf) == "\n" {
			fmt.Printf("-")
		} else {
			fmt.Printf("%v", string(buf))
		}
		count++
	}

	fmt.Println("Done.")

}
