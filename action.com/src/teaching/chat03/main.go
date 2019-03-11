package main

import "fmt"

//创建字典

type Dictionary struct {
	data map[interface{}]interface{}
}

func (dic *Dictionary) Get(key interface{}) interface{} {
	return dic.data[key]
}

func (dic *Dictionary) Set(key, value interface{}) {
	dic.data[key] = value
}

func (dic *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range dic.data {
		if !callback(k, v) {
			return
		}
	}
}

func (dic *Dictionary) Clear() {
	dic.data = make(map[interface{}]interface{})
}

func NewDictionary() *Dictionary {
	dic := &Dictionary{}
	dic.Clear()
	return dic
}

func main() {

	dict := NewDictionary()

	dict.Set("MyLove", 60)
	dict.Set("MyHeart", 60)
	dict.Set("MyEngine", 60)

	love := dict.Get("MyLove")
	fmt.Printf("love:%d\n", love)

	dict.Visit(func(k, v interface{}) bool {
		fmt.Printf("key : %s value : %d \n", k, v)
		return true
	})

}
