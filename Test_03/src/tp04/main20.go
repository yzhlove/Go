package main

import "fmt"

// 切片

func main() {

	x := [...]int{0,1,2,3,4,5,6,7,8,9}

	a := x[:]
	b := x[2:5]
	c := x[2:5:7]
	d := x[4:]
	e := x[:4]
	f := x[:4:6]

	fmt.Printf("%+v len = %d,cap = %d \n",a,len(a),cap(a))
	fmt.Printf("%+v len = %d,cap = %d \n",b,len(b),cap(b))
	fmt.Printf("%+v len = %d,cap = %d \n",c,len(c),cap(c))
	fmt.Printf("%+v len = %d,cap = %d \n",d,len(d),cap(d))
	fmt.Printf("%+v len = %d,cap = %d \n",e,len(e),cap(e))
	fmt.Printf("%+v len = %d,cap = %d \n",f,len(f),cap(f))

}

/*
[0 1 2 3 4 5 6 7 8 9] len = 10,cap = 10
[2 3 4] len = 3,cap = 8
[2 3 4] len = 3,cap = 5
[4 5 6 7 8 9] len = 6,cap = 6
[0 1 2 3] len = 4,cap = 10
[0 1 2 3] len = 4,cap = 6
*/