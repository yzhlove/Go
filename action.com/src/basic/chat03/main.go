package main

import (
	"bytes"
	"fmt"
)

//字符串高效拼接

func main() {

	var stringBuild bytes.Buffer

	str := "apple "

	for i := 0; i < 10; i++ {
		stringBuild.WriteString(str)
	}

	fmt.Printf("str = %s \n", stringBuild.String())

}
