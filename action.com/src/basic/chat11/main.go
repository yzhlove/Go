package main

import (
	"fmt"
	"strings"
)

//链式处理

func main() {

	list := []string{
		"go scanner",
		"go parse",
		"go compiler",
		"go printer",
		"go formater",
	}

	funcs := []func(string) string{
		delPrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	stringProcess(list, funcs)

	for _, value := range list {
		fmt.Println("str = ", value)
	}

}

func delPrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func stringProcess(list []string, funcs []func(string) string) {

	for index, str := range list {
		result := str
		for _, fn := range funcs {
			result = fn(result)
		}
		list[index] = result
	}
}
