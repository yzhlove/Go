package main

import "fmt"

//copy使用

func main() {

	s1 :=[]int{1,2,3,4,5}
	s2 :=[]int{6,7,8,9,10,11,12}

	copy(s2,s1)
	//s1 = [1 2 3 4 5] s2 = [1 2 3 4 5 11 12]
	fmt.Printf("s1 = %v s2 = %v \n",s1,s2)

	t1 := []int{3,4,5,6,7,8}
	t2 := []int{1,2,3}
	copy(t2,t1)
	//t1 = [3 4 5 6 7 8] t2 = [3 4 5]
	fmt.Printf("t1 = %v t2 = %v \n",t1,t2)


	m1 := []int{1,2,3}
	var m2 []int
	copy(m2,m1)
	//m1 = [1 2 3] m2 = []
	fmt.Printf("m1 = %v m2 = %v \n",m1,m2)

	w1 := []int{1,2,3,4,5,6}
	w2 := make([]int,0,16)
	copy(w2,w1)
	//w1 = [1 2 3 4 5 6] w2 = []
	fmt.Printf("w1 = %v w2 = %v \n",w1,w2)

	x1 := []int{1,2,3,4,5,6}
	x2 := make([]int,8,16)
	//x1 = [1 2 3 4 5 6] x2 = [1 2 3 4 5 6 0 0]
	copy(x2,x1)
	fmt.Printf("x1 = %v x2 = %v \n",x1,x2)

}
