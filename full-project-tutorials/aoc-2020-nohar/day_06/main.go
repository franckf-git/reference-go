package main

import (
	"fmt"
	"math"
	"math/bits"

	"gitlab.com/neuware/aoc-2020/utils"
)

func main() {
	var part1, part2 int
	input := utils.ReadLines()

	fmt.Println("--- Using maps ---")
	part1, part2 = solveWithMaps(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("--- Using slices ---")
	part1, part2 = solveWithSlices(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("--- Using bitwise ops ---")
	part1, part2 = solveBitwise(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

// Naive solution (using maps)

type Answers map[rune]int

func iterAnswers(input []string, cbk func(Answers, int)) {
	answers := make(Answers) // group answers
	size := 0                // group size
	for _, line := range input {
		if line == "" {
			cbk(answers, size)
			answers = make(Answers)
			size = 0
			continue
		}

		size++
		for _, c := range line {
			answers[c]++
		}
	}
	cbk(answers, size)
}

func solveWithMaps(input []string) (int, int) {
	var part1, part2 int
	iterAnswers(input, func(a Answers, size int) {
		part1 += len(a)
		for _, n := range a {
			if n == size {
				part2++
			}
		}
	})
	return part1, part2
}

// A faster solution, with comparable semantics, but only 1 memory allocation

var zeros [26]int

func iterAnswerSlices(input []string, cbk func([]int, int)) {
	answers := make([]int, 26) // group answers
	size := 0                  // group size
	for _, line := range input {
		if line == "" {
			cbk(answers, size)
			copy(answers, zeros[:])
			size = 0
			continue
		}

		size++
		for _, c := range line {
			answers[byte(c)-'a']++
		}
	}
	cbk(answers, size)
}

func solveWithSlices(input []string) (int, int) {
	var part1, part2 int
	iterAnswerSlices(input, func(a []int, size int) {
		for _, n := range a {
			if n > 0 {
				part1++
			}
			if n == size {
				part2++
			}
		}
	})
	return part1, part2
}

// A much more efficient solution using bitwise operations

const (
	initUnion = 0
	initInter = math.MaxUint32
)

func solveBitwise(input []string) (int, int) {
	var part1, part2 int
	var union, inter uint32 = initUnion, initInter
	for _, line := range input {
		if line == "" {
			part1 += bits.OnesCount32(union)
			part2 += bits.OnesCount32(inter)
			union, inter = initUnion, initInter
			continue
		}
		var answers uint32
		for _, c := range line {
			answers |= 1 << (byte(c) - 'a')
		}
		union |= answers
		inter &= answers
	}
	// Flush the last group
	part1 += bits.OnesCount32(union)
	part2 += bits.OnesCount32(inter)
	return part1, part2
}
