package main

import (
	"testing"
)

func TestPart2(t *testing.T) {
	cases := []struct {
		Input    []int
		Expected int
	}{
		{
			[]int{17, 0, 13, 19},
			3417,
		},
		{
			[]int{7, 13, 0, 0, 59, 0, 31, 19},
			1068781,
		},
		{
			[]int{1789, 37, 47, 1889},
			1202161486,
		},
	}

	for i, c := range cases {
		res := part2(c.Input)
		if res != c.Expected {
			t.Fatalf("case %d: expected %d, got %d", i+1, c.Expected, res)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	_, buses := parseInput()
	for n := 0; n < b.N; n++ {
		if part2(buses) == 0 {
			b.Fail()
		}
	}
}
