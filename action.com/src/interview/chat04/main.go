package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("People ShowA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("People ShowB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("Teacher ShowB")
}

func main() {
	t := &Teacher{}
	t.ShowA()
}
