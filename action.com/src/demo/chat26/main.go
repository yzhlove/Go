package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//斐波拉切数列改版

type intGen func() int

func fib() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func (g intGen) Read(b []byte) (int, error) {

	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	str := fmt.Sprintf("%d\n", next)
	return strings.NewReader(str).Read(b)
}

func printFileContext(reader io.Reader) {

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {

	fb := fib()
	printFileContext(fb)

}
