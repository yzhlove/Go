package main

import (
	"bytes"
	"fmt"
)

//可变参数列表

func printType(slist ...interface{}) string {
	var b bytes.Buffer
	for _, s := range slist {
		str := fmt.Sprintf("%v", s)
		typeString := ""
		switch s.(type) {
		case int:
			typeString = "int"
		case string:
			typeString = "string"
		case bool:
			typeString = "bool"
		}
		b.WriteString(typeString)
		b.WriteString(":")
		b.WriteString(str)
		b.WriteString("\n")
	}
	return b.String()
}

func main() {

	temp := []interface{}{1, 2, 3, "love", "xjj", "lcm", "xyj", "hxy", true, false}
	res := printType(temp...)
	fmt.Println(res)
}
