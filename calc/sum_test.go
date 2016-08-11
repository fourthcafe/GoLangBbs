package calc

import "testing"

type TestData struct {
	arg1   int
	arg2   int
	result int
}

var testData = []TestData{
	{2, 6, 8},
	{-8, 3, -5},
	{6, -6, 0},
	{0, 0, 0},
}

func TestSum(t *testing.T) {
	for _, data := range testData {
		sumResult := Sum(data.arg1, data.arg2)
		if sumResult != data.result {
			t.Errorf("%d + %d의 결과값이 %d가 아님\nresult:%d", data.arg1, data.arg2, data.result, sumResult)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(3, 4)
	}
}
