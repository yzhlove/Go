package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "yzh", Age: 24},
		{Name: "lcm", Age: 23},
		{Name: "xjj", Age: 22},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	fmt.Printf("%v \n", m)
	for n, v := range m {
		fmt.Printf("%s %v", n, v)
	}
}

func main() {
	pase_student()
}
