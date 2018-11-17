package main

import (
	"fmt"
	"strconv"
)

//map函数之间以引用传递

type HashMap map[string]string

var hashMap HashMap

func init() {
	hashMap = make(map[string]string)
}

func (hmp HashMap)show() {
	for index,value :=range hmp {
		fmt.Printf("key:%v value:%v \n",index,value)
	}
}

//hashMap按引用传递，无需传图指针
func (hmp HashMap) remove(key string) {
	fmt.Printf("==== delete %v ==== \n",hmp[key])
	delete(hmp,key)
}

func main() {

	colors := []string{"Red","Green","Yellow","Blue","Gary"}

	for index,value := range colors {
		hashMap[strconv.Itoa(index)] = value
	}

	hashMap.show()

	hashMap.remove("1")

	hashMap.show()

}