package main

// 引用类型使用

func mkslice() []int {

	arr := make([]int,0,10)
	arr = append(arr,100)
	return arr
}

func mkmap() map[string]int {
	mmp := make(map[string]int)
	mmp["a"] = 1
	return mmp
}

func main() {

	mmp := mkmap()
	print(mmp["a"])

	println("-------------")

	arr := mkslice()
	print(arr[0])

}