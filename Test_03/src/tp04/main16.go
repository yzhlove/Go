package main

// 内置函数 len 和 cap都返回第一为度的长度

func main() {

	a := [2]int{}
	b := [...][2]int{
		{10,20},
		{20,30},
		{30,40},
	}

	println(len(a),cap(a))
	println(len(b),cap(b))
	println(len(b[1]),cap(b[1]))

}
