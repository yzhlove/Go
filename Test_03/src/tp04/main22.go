package main

// 切片对象的创建

func main() {

	s1 := make([]int,3,5)	// len =3 cap = 5 所有元素初始化为0
	s2 := make([]int,7)		// len = cap = 7 ，所有元素初始化为0
	s3 := []int{10,20,5:30} // len = cap = 6 ，index = 5的元素值为30

	println("s1 = ",s1,len(s1),cap(s1))
	println("s2 = ",s2,len(s2),cap(s2))
	println("s3 = ",s3,len(s3),cap(s3))

}
