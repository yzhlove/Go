package main

import (
	"reflect"
	"testing"
)

type Data struct {
	Hp int
}

func Benchmark_NativeAssign(b *testing.B) {
	v := Data{Hp: 2}
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Hp = 3
	}
}

func Benchmark_ReflectAssign(b *testing.B) {
	v := Data{Hp: 2}
	refv := reflect.ValueOf(&v).Elem()
	f := refv.FieldByName("Hp")
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f.SetInt(3)
	}
}

func Benchmark_ReflectAssign2(b *testing.B) {
	v := Data{Hp: 2}
	refv := reflect.ValueOf(&v).Elem()
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		refv.FieldByName("Hp").SetInt(3)
	}
}
