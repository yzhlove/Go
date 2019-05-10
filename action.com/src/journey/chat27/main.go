package main

import "fmt"

//接口的嵌套与组合

type Writer interface {
	write(b []byte) (int, error)
}

type Reader interface {
	read() ([]byte, error)
}

type ReadWrite interface {
	Writer
	Reader
}

type Device struct {
	data string
}

func (d *Device) write(b []byte) (int, error) {
	d.data = string(b)
	fmt.Println("write string:", d.data)
	return len(d.data), nil
}

func (d *Device) read() ([]byte, error) {
	return []byte(d.data), nil
}

func main() {

	var (
		data   []byte
		number int
		err    error
	)

	device := new(Device)
	device.data = "what are you doing"

	var rw ReadWrite = device

	if data, err = rw.read(); err != nil {
		panic(err)
	}
	fmt.Println("read data = ", string(data))

	if number, err = rw.write([]byte("Are you ok")); err != nil {
		panic(err)
	}
	fmt.Println("write number = ", number)

	var onlyRead Reader = device

	if data, err = onlyRead.read(); err != nil {
		panic(err)
	}

	fmt.Println("onlyRead = ", string(data))

}
