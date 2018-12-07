package main

import "fmt"

//访问者模式

//定义访问者接口
type vaster interface {
	Visit()
}

//生产环境
type productVaster struct {}

//测试环境
type testVaster struct {}

func (pv *productVaster) Visit() {
	fmt.Println("This is product...")
}

func (tv *testVaster) Visit() {
	fmt.Println("This is test...")
}

//定义元素接口
type element struct {}

func (e *element) Accept(v vaster) {
	v.Visit()
}

//修改
type envExample struct {
	element
}

func (e *envExample) Print(v vaster) {
	e.element.Accept(v)
}

//使用
func main() {

	e := new(element)
	e.Accept(new(productVaster))
	e.Accept(new(testVaster))

	m := new(envExample)
	m.Print(new(productVaster))
	m.Print(new(testVaster))

}
