package main

import (
	"testing"

	"gitlab.com/neuware/aoc-2020/utils"
)

var testInput = utils.ReadLines()

func BenchmarkWithMaps(b *testing.B) {
	var part1, part2 int
	for n := 0; n < b.N; n++ {
		part1, part2 = solveWithMaps(testInput)
		if part1 < part2 {
			b.Errorf("wrong result")
		}
	}
}

func BenchmarkWithSlices(b *testing.B) {
	var part1, part2 int
	for n := 0; n < b.N; n++ {
		part1, part2 = solveWithSlices(testInput)
		if part1 < part2 {
			b.Errorf("wrong result")
		}
	}
}

func BenchmarkBitwise(b *testing.B) {
	var part1, part2 int
	for n := 0; n < b.N; n++ {
		part1, part2 = solveBitwise(testInput)
		if part1 < part2 {
			b.Errorf("wrong result")
		}
	}
}
