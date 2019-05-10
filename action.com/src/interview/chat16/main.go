package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string {
	return "不允许使用带有递归调用的方法!"
}

func main() {

	p := &People{}
	fmt.Println(p)

}
