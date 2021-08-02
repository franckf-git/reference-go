package main

import (
	"sort"
	"testing"

	"gitlab.com/neuware/aoc-2020/utils"
)

func TestArrangements(t *testing.T) {
	cases := []struct {
		Input    []int
		Expected int
	}{
		{
			[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4},
			8,
		},
		{
			[]int{
				28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49,
				45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4,
				2, 34, 10, 3,
			},
			19208,
		},
	}
	for i, c := range cases {
		sort.Ints(c.Input)
		res := arrangements(c.Input)
		if res != c.Expected {
			t.Fatalf("case %d, expected %d, got %d", i, c.Expected, res)
		}
	}
}

func BenchmarkArrangements(b *testing.B) {
	input := utils.ReadInts()
	sort.Ints(input)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if arrangements(input) == 0 {
			b.Fail()
		}
	}
}
