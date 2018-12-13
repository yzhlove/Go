package main

import (
	"container/list"
	"fmt"
)

//list 双链表

func main() {

	list := list.New()

	//尾部添加
	list.PushBack("yzh")

	//头部添加
	list.PushFront(123)

	//尾部添加之后保存句柄
	element := list.PushBack("xjj")

	//在element之后添加元素
	list.InsertAfter("xyj", element)

	//在element之前添加元素
	list.InsertBefore("lcm", element)

	for index := list.Front(); index != nil; index = index.Next() {
		fmt.Printf("value: %v \n", index.Value)
	}

	fmt.Println()

	//删除特定元素
	list.Remove(element)

	for index := list.Front(); index != nil; index = index.Next() {
		fmt.Printf("value: %v \n", index.Value)
	}

}
