package main

import "fmt"

//切片思考

func main() {

	arr := [...]int{0,1,2,3,4,5,6,7}

	s1 :=arr[2:6]
	s2 :=s1[3:5]

	//[2 3 4 5] [5 6]
	fmt.Printf("%v %v \n",s1,s2)

	s3 := s1[3:6]
	//s4 := s1[3:7]
	//[5 6 7]
	fmt.Printf("%v %v \n",s3)


	t1 := arr[2:6:6]
	t2 := t1[3:5]
	//slice bounds out of range cap（ 被限定）
	fmt.Printf("%v %v \n",t1,t2)


}
