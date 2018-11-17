package main

import "fmt"

//当切片的长度大于总容量长度时，使用append会生成一个新的切片

func main() {


	//创建一个长度为0，容量为5的切片
	slice := make([]string,0,5)

	str := [5]string{"A","B","C","D","E"}

	for index,_ := range str {
		//fmt.Printf("%+v %+v \n",index,str[index])
		//slice[index] = str[index]
		slice = append(slice,str[index])
	}

	//重新分配一个切片
	// len :3-2
	// cap :4-2
	slice2 := slice[2:3:4]

	slice2 = append(slice2,[]string{"E","F","G"}...)

	//append函数:
	//当新加入的元素大于切片的容量的时候，append函数会重新开辟一块新的内存

	fmt.Printf("slice:%#v slice2:%#v \n",slice,slice2)
	fmt.Printf("slice len:%v cap %v slice2 len:%v cap:%v \n",len(slice),cap(slice),len(slice2),cap(slice2))

}
