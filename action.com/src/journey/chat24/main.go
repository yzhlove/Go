package main

import (
	"fmt"
	"sort"
)

//自定义排序

type MyString []string

func (ms MyString) Len() int {
	return len(ms)
}

func (ms MyString) Less(i, j int) bool {
	return ms[i] < ms[j]
}

func (ms MyString) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func main() {

	names := MyString{
		"1.yzh",
		"5.cjj",
		"4.lcm",
		"2.xyj",
		"3.xjj",
	}

	sort.Sort(names)

	for _, v := range names {
		fmt.Printf("%s \n", v)
	}

}
