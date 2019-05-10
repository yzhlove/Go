package main

import (
	"fmt"
	"regexp"
)

//正则表达式

const text = `
my  name is ccmouse@email.com@456.com
email.com is lcmm5201314@gmail.com
email is lcm_520@live.com
email is 1228153231@qq.com
sina is sina@mail.com.cn
`

func main() {
	var (
		reg *regexp.Regexp
	)
	reg = regexp.MustCompile(`([a-zA-Z0-9_]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)
	matchs := reg.FindAllStringSubmatch(text, -1)

	for _, vs := range matchs {
		for _, v := range vs {
			fmt.Printf("%s ", v)
		}
		fmt.Println()
	}

}
