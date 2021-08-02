package main

import "testing"

func TestMemoryGame(t *testing.T) {
	cases := []struct {
		Input    []int32
		Expected int32
	}{
		{Input: []int32{0, 3, 6}, Expected: 436},
		{Input: []int32{1, 3, 2}, Expected: 1},
		{Input: []int32{2, 1, 3}, Expected: 10},
	}
	for _, c := range cases {
		res := MemoryGame(c.Input, 2020)
		if res != c.Expected {
			t.Fatalf("case %v, expected %d, got %d", c.Input, c.Expected, res)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if MemoryGame(Input, 2020) != 959 {
			b.Fail()
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if MemoryGame(Input, 30000000) != 116590 {
			b.Fail()
		}
	}
}

func BenchmarkPart1Hashmap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if MemoryGameHashmap(Input, 2020) != 959 {
			b.Fail()
		}
	}
}

func BenchmarkPart2Hashmap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if MemoryGameHashmap(Input, 30000000) != 116590 {
			b.Fail()
		}
	}
}
