package main

import (
	"testing"

	"gitlab.com/neuware/aoc-2020/utils"
)

func BenchmarkRun(b *testing.B) {
	program := load(utils.ReadLines())
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		res := run(program)
		if res == -1 {
			b.Fail()
		}
	}
}

func BenchmarkPatch(b *testing.B) {
	program := load(utils.ReadLines())
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		res := patch(program)
		if res == -1 {
			b.Fail()
		}
	}
}

func BenchmarkBruteForce(b *testing.B) {
	program := load(utils.ReadLines())
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		res := bruteForce(program)
		if res == -1 {
			b.Fail()
		}
	}
}
