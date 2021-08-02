package main

import (
	"fmt"

	"gitlab.com/neuware/aoc-2020/utils"
)

// This is the default, "naive" implementation of the algorithm.
func seatID(code string) int {
	minRow, maxRow, minSeat, maxSeat := 0, 127, 0, 7
	for _, c := range code {
		switch c {
		case 'F':
			maxRow -= (maxRow - minRow + 1) / 2
		case 'B':
			minRow += (maxRow - minRow + 1) / 2
		case 'L':
			maxSeat -= (maxSeat - minSeat + 1) / 2
		case 'R':
			minSeat += (maxSeat - minSeat + 1) / 2
		}
	}
	return minRow*8 + minSeat
}

// This implementation is based on the fact that seat codes are
// just a binary representation of the seat IDs:
// * 'L' and 'F' map to the digit 0
// * 'B' and 'R' map to the digit 1
// Once this is figured out, converting a code to a seat ID is simply parsing
// the binary representation of a number.
func seatID2(code string) int {
	var res int
	for _, c := range code {
		res <<= 1
		if c == 'B' || c == 'R' {
			res |= 1
		}
	}
	return res
}

func main() {
	min, max := 1024, 0
	checked := make([]bool, 1024)

	for _, code := range utils.ReadLines() {
		id := seatID(code)
		if id > max {
			max = id
		}
		if id < min {
			min = id
		}
		checked[id] = true
	}

	fmt.Println("Part 1:", max)
	for i := min; i <= max; i++ {
		if !checked[i] {
			fmt.Println("Part 2:", i)
			return
		}
	}
}
