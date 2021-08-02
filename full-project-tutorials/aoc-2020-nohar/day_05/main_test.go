package main

import (
	"testing"

	"gitlab.com/neuware/aoc-2020/utils"
)

var cases = []struct {
	Input    string
	Expected int
}{
	{"BFFFBBFRRR", 567},
	{"FFFBBBFRRR", 119},
	{"BBFFBBFRLL", 820},
}

var testInput = utils.ReadLines()

func TestSeatID(t *testing.T) {
	for _, c := range cases {
		id := seatID(c.Input)
		if id != c.Expected {
			t.Errorf("Input: %q, Expected: %d, Got: %d", c.Input, c.Expected, id)
		}
	}
}

func TestSeatID2(t *testing.T) {
	for _, c := range cases {
		id := seatID2(c.Input)
		if id != c.Expected {
			t.Errorf("Input: %q, Expected: %d, Got: %d", c.Input, c.Expected, id)
		}
	}
}

func BenchmarkSeatID(b *testing.B) {
	L := len(testInput)
	for n := 0; n < b.N; n++ {
		if seatID(testInput[n%L]) == 0 {
			b.Errorf("wrong result")
		}
	}
}

func BenchmarkSeatID2(b *testing.B) {
	L := len(testInput)
	for n := 0; n < b.N; n++ {
		if seatID2(testInput[n%L]) == 0 {
			b.Errorf("wrong result")
		}
	}
}
