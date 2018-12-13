package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//拼接字符串

func main() {

	var b bytes.Buffer

	b.Write([]byte("Hello"))

	fmt.Fprintf(&b, " world")

	io.Copy(os.Stdout, &b)

	fmt.Printf("%T %v] \n", b, b)

}
