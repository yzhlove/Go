package main

import (
	"fmt"
)

type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

type welcome string

func (w welcome) selfInfo(k, v interface{}) {
	fmt.Printf("%s,我叫%s ,今年%d \n", w, k, v)
}

func Each(m map[interface{}]interface{}, h HandlerFunc) {
	if m == nil || len(m) == 0 {
		return
	}
	for k, v := range m {
		h.Do(k, v)
	}
}

func main() {

	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 21
	persons["王五"] = 22
	persons["赵六"] = 21

	var w welcome = "Hello "
	Each(persons, HandlerFunc(w.selfInfo))

}
