package utils

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func BenchmarkAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Average([]float64{1, 2})
	}
}

// go test -bench=.
