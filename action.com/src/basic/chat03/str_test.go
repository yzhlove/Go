package main

import (
	"bytes"
	"testing"
)

const str = "what are you doing . "

func BenchmarkAppendStr(b *testing.B) {

	var stringBuild bytes.Buffer

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stringBuild.WriteString(str)
	}

}

func BenchmarkAddStr(b *testing.B) {
	var temp string
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		temp += str
	}
}
