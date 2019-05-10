package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

//bufio.NewReader 测试

func main() {

	inputReader := strings.NewReader("1234567890")
	bufferReader := bufio.NewReader(inputReader)
	Read(*bufferReader)
	buffer2Reader := bufio.NewReader(inputReader)
	Read(*buffer2Reader)
	_, _ = bufferReader.Peek(3)
	Read(*bufferReader)
}

func Read(reader bufio.Reader) {
	bytes, err := ioutil.ReadAll(&reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("reader:%s \n", bytes)
}
