package chat01

import (
	"math"
	"strings"
	"testing"
)

//一个简单的表格驱动测试

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func TestTriangle(t *testing.T) {

	testTable := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, temp := range testTable {
		if result := calcTriangle(temp.a, temp.b); result != temp.c {
			t.Errorf("calc Err:%d %d %d success:%d ", temp.a, temp.b, temp.c, result)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {

	str := "黑化肥挥发发灰会花飞灰化肥挥发非黑会飞花"
	chars := []rune(str)
	for _, v := range chars {
		if strings.Index(str, string(v)) == -1 {
			b.Errorf("find Err:%c ", v)
		}
	}
}

//性能测试案例
func BenchmarkSearch(b *testing.B) {

	value := make(map[int]int)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			for k := 0; k < 100; k++ {
				value[i] += j * k
			}
		}
	}

	for k, v := range value {
		b.Logf("index:%d value:%d", k, v)
	}
}
