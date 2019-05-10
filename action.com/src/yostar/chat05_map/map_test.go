package main

import "testing"

func Benchmark_AutoMap(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	data := make(map[int]int, 10000)
	for i := 0; i < b.N; i++ {
		data[i] = i
	}

}

func Benchmark_Map(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	data := make(map[int]int)
	for i := 0; i < b.N; i++ {
		data[i] = i
	}
}
