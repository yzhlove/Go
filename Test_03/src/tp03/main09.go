package main

// 匿名函数与结构体

func testStruct() {

	type calc struct {
		mul func(x,y int) int
	}

	x := calc{
		mul:func (x,y int) int {
			return (x << 1) + (y << 1)
		},
	}

	println(x.mul(2,3))

}


func testChannel() {

	c := make(chan func(int,int) int ,2)
	c <- func(i int, i2 int) int {
		return i + i2
	}

	println((<-c)(1,2))

}

func main() {

	testStruct()
	testChannel()


}