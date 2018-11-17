package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	//线程安全

	//方式一
	var syncMap sync.Map

	//方式二
	//syncMap := new(sync.Map)

	colors := [5]string{"Red","Orange","Yellow","Green","Blue"}

	for index,value := range colors {
		syncMap.Store(index,value)
	}

	//便利
	syncMap.Range(func(k,v interface{}) bool {
		fmt.Printf("%v , %v \n",k,v)
		return true
	})

	//根据key取value
	value, exists := syncMap.Load(rand.Intn(5) + 1)
	if exists {
		fmt.Printf("Load Value: %v \n",value)
	}

	//根据key删除对应的value
	syncMap.Delete(rand.Intn(5) + 1)

	syncMap.Range(func(k,v interface{}) bool {
		fmt.Printf("%v , %v \n",k,v)
		return true
	})


}
