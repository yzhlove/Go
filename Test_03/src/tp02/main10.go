package main


// 如果range对象是函数，则仅仅被执行一次

var count = 1

func dataValue() []int {
	 println("count = ",count)
	 count++
	 return []int{10,20,30}
}

func  main() {

	for i ,x := range dataValue() {
		println("i = ",i, " x = ",x)
	}

}

/*
count =  1
i =  0  x =  10
i =  1  x =  20
i =  2  x =  30
*/