package main

import (
	"fmt"
	"reflect"
)

func add(a, b int) int {
	return a + b
}

func main() {
	//将函数包装为反射对象
	funcValue := reflect.ValueOf(add)
	//传参数
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	//调用
	retList := funcValue.Call(paramList)
	//获取返回值
	fmt.Println(retList[0].Int())
}
