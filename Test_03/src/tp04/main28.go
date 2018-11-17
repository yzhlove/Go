package main

import "fmt"

// 向切片尾部添加数据

//超出cap的限制则回重新非配内存

func main() {

	s := make([]int,0,5)
	s1 := append(s,10)
	s2 := append(s,20,30,40)
	s3 := append(s,12,3,4,5,6,7,8)
	fmt.Println(s , len(s),cap(s))
	fmt.Println(s1 , len(s1),cap(s1))
	fmt.Println(s2 , len(s2),cap(s2))
	fmt.Println(s3 , len(s3),cap(s3))

}

/*
[] 0 5
[20] 1 5
[20 30 40] 3 5
[12 3 4 5 6 7 8] 7 10
 */