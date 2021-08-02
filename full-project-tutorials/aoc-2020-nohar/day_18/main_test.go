package main

import (
	"testing"

	"gitlab.com/neuware/aoc-2020/utils"
)

func TestEvalPart1(t *testing.T) {
	cases := []struct {
		Input    string
		Expected int
	}{
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, c := range cases {
		res := evalPart1(c.Input)
		if res != c.Expected {
			t.Errorf("Case %q, expected %d, got %d", c.Input, c.Expected, res)
		}
	}
}

func TestEvalPart2(t *testing.T) {
	cases := []struct {
		Input    string
		Expected int
	}{
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"2 + 4 * 9", 54},
		{"6 + 9 * 8 + 6", 210},
		{"(6 + 9 * 8 + 6) + 6", 216},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	for _, c := range cases {
		res := evalPart2(c.Input)
		if res != c.Expected {
			t.Errorf("Case %q, expected %d, got %d", c.Input, c.Expected, res)
		}
	}
}

func BenchmarkEvalPart1(b *testing.B) {
	input := utils.ReadLines()
	size := len(input)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if evalPart1(input[n%size]) < 0 {
			b.Fail()
		}
	}
}

func BenchmarkEvalPart2(b *testing.B) {
	input := utils.ReadLines()
	size := len(input)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if evalPart2(input[n%size]) < 0 {
			b.Fail()
		}
	}
}
