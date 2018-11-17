package main

import "fmt"

//   定义多位数组的时候，仅允许第一纬度使用 ...

func main () {

	a := [2][2]int {
		{1,3},
		{2,4},
	}

	b := [...][2]int{
		{10,20},
		{30,40},
	}

	c := [...][2][2]int{
		{
			{1,2},
			{3,4},
		},{
			{10,20},
			{30,40},
		},
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%+v\n",a)
	fmt.Printf("%+v\n",b)
	fmt.Printf("%+v\n",c)

}