package main

import (
	"fmt"

	"gitlab.com/neuware/aoc-2020/utils"
)

func main() {
	input := utils.ReadLines()
	grid, count := run(input, CanTakeSeatPart1, MustLeaveSeatPart1)

	fmt.Println("finished after", count, "generations")
	fmt.Println("Part 1:", grid.Count(Occupied))

	grid, count = run(input, CanTakeSeatPart2, MustLeaveSeatPart2)
	fmt.Println("finished after", count, "generations")
	fmt.Println("Part 2:", grid.Count(Occupied))
}

type SeatEvalFunc func(*Grid, int, int) bool

func run(input []string, canTake, mustLeave SeatEvalFunc) (*Grid, int) {
	src := NewGridFromInput(input)
	dst := NewGrid(src.Width(), src.Height())

	var count int
	changes := 1
	for changes != 0 {
		changes = step(src, dst, canTake, mustLeave)
		src, dst = dst, src
		count++
	}
	return src, count
}

func step(src, dst *Grid, canTake, mustLeave SeatEvalFunc) int {
	var changes int
	for y := 0; y < src.Height(); y++ {
		for x := 0; x < src.Width(); x++ {
			switch src.At(x, y) {
			case Empty:
				if canTake(src, x, y) {
					dst.Set(x, y, Occupied)
					changes++
					break
				}
				dst.Set(x, y, Empty)
			case Occupied:
				if mustLeave(src, x, y) {
					dst.Set(x, y, Empty)
					changes++
					break
				}
				dst.Set(x, y, Occupied)
			}
		}
	}
	return changes
}
