package basic

import (
	"math"
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
