package main

import "testing"

func Benchmark_SliceInit(t *testing.B) {

	t.StopTimer()
	t.ResetTimer()
	t.StartTimer()

	var array []int

	for i := 0; i < t.N; i++ {
		array = append(array, i)
	}

}

func Benchmark_SliceMalloc(t *testing.B) {
	t.StopTimer()
	t.ResetTimer()
	t.StartTimer()

	length := t.N
	array := make([]int, 0, length)

	for i := 0; i < length; i++ {
		array = append(array, i)
	}

}
