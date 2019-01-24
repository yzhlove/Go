package main

import "fmt"

//一个简单的接口

type DataWriter interface {
	write(data interface{}) error
}

type file struct{}

func (f *file) write(data interface{}) error {
	fmt.Println("fwrite:", data)
	return nil
}

func main() {
	f := new(file)
	var dw DataWriter
	dw = f
	_ = dw.write(map[string]string{"Hello": "_world"})

}
