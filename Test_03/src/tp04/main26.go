package main

import "testing"

// 切片内存占比

func array() [1024]int {
	var x [1024]int
	for i := 0;i < len(x);i++ {
		x[i] = i
	}
	return x
}

func slice() []int {
	x := make([]int,1024)
	for i := 0;i < len(x);i++ {
		x[i] = i
	}
	return x
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N;i++ {
		array()
	}
}


func BenchmarkSlice(b *testing.B) {
	for i := 0;i < b.N;i++ {
		slice()
	}
}


func main() {



}
