package main

import (
	"fmt"

	"gitlab.com/neuware/aoc-2020/utils"
)

func findTarget(input []int) int {
SearchLoop:
	for i := 25; i < len(input); i++ {

		// For each j, k such that: i-25 < j < k < i
		for j := i - 25; j < i; j++ {
			for k := j + 1; k < i; k++ {

				if input[i] == input[j]+input[k] {
					// We've found the right (j, k) pair,
					// move on to next i.
					continue SearchLoop
				}
			}
		}

		// If we reach here, then i doesn't have j, k that
		// satisfies the constraint: this is our target
		return input[i]
	}
	panic("didn't find target")
}

func findWeakness(target int, input []int) int {
	for i := 0; i < len(input); i++ {
		sum, min, max := 0, input[i], input[i]
		for j := i + 1; j < len(input); j++ {
			x := input[j]
			sum += x
			if x < min {
				min = x
			}
			if x > max {
				max = x
			}
			if sum == target {
				return min + max
			}
			if sum > target {
				break
			}
		}
	}
	panic("didn't find weakness")
}

func main() {
	input := utils.ReadInts()

	target := findTarget(input)
	fmt.Println("Part 1:", target)
	fmt.Println("Part 2:", findWeakness(target, input))

}
