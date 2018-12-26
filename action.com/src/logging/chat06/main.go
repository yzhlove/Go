package main

import (
	"bytes"
	"fmt"
	"os"
)

//io.Write的使用

func main() {
	//一个Hello World
	var b bytes.Buffer
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "World")
	b.WriteTo(os.Stdout)
}
