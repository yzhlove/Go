package main

import "fmt"

// 修改map的正确做法

type user struct {

	name string
	age byte
}


func main() {


	m := map[int]user{
		1 : {"Tom",19},
	}

	//m[1].age += 1		// ERROR
	// 第一种修改
	temp := m[1]
	temp.age += 1
	m[1] = temp

	fmt.Printf("%+v \n",m)

	m2 := map[int]*user {
		1 : &user{"Jack",20},
	}
	m2[1].age++
	fmt.Printf("%+v \n",*m2[1])

}
