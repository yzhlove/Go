package main

// 复合类型的初始化

type data struct {
	name string
	age int
}

var da = data{"hello ",23}
var db = data{
	"world",
	45,		//如果换行 必须以","结尾
}
//dc := []int {
//	1,2,
//	3,
//	4,5,
//}


