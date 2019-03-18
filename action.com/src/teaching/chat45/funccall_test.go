package main

import (
	"reflect"
	"testing"
)

func foo(v int) {}

func Benchmark_NativeCall(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		foo(i)
	}
}

func Benchmark_reflectCall(b *testing.B) {
	v := reflect.ValueOf(foo)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Call([]reflect.Value{reflect.ValueOf(2)})
	}
}
