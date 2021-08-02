package main

import (
	"fmt"
	"sort"

	"gitlab.com/neuware/aoc-2020/utils"
)

func arrangements(input []int) int {
	x := 0                       // Last computed joltage (initialize on the 0j input)
	sum, prev1, prev2 := 1, 0, 0 // Number of arrangements for x, x-1, x-2 jolts
	for _, n := range input {
		for x < n-1 {
			// "Slide the window" until we have the number of arrangements for
			// n-1, n-2, n-3 jolts: fill missing values with 0.
			x++
			prev2, prev1, sum = prev1, sum, 0
		}
		x++
		prev2, prev1, sum = prev1, sum, sum+prev1+prev2
	}
	return sum
}

func diffs(input []int) int {
	var ones, threes int
	x := 0 // count the 0j input
	for _, y := range input {
		switch y - x {
		case 1:
			ones++
		case 3:
			threes++
		}
		x = y
	}
	threes++ // count the max+3j device
	return ones * threes
}

func main() {
	input := utils.ReadInts()
	sort.Ints(input)

	fmt.Println("Part 1:", diffs(input))
	fmt.Println("Part 2:", arrangements(input))
}
