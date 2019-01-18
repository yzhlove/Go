package main

import "fmt"

//声明一个解析错误

type ParseError struct {
	Filename string
	Line     int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Filename:%s :Line%d \n", e.Filename, e.Line)
}

func newParseError(filename string, line int) error {
	return &ParseError{Filename: filename, Line: line}
}

func main() {

	var e error
	e = newParseError("main.go", 1)

	fmt.Println(e.Error())

	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("filename:%s line:%d \n", detail.Filename, detail.Line)
	default:
		fmt.Println("other error")
	}
}
