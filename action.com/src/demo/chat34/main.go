package main

import (
	"fmt"
)

//接口示例

type Handler interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m == nil || len(m) == 0 {
		return
	}
	for k, v := range m {
		h.Do(k, v)
	}
}

type welcome string

func (w welcome) Do(k, v interface{}) {
	fmt.Printf("%s , 我叫%s ，今年%d \n", w, k, v)
}

func main() {

	person := make(map[interface{}]interface{})
	person["张三"] = 20
	person["李四"] = 21
	person["王五"] = 22
	person["赵六"] = 23

	var w welcome = " Hello !"

	Each(person, w)

}
