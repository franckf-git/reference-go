package main

import (
	"fmt"

	"gitlab.com/neuware/aoc-2020/utils"
)

func main() {
	input := utils.ReadLines()
	world3D := World3DFromInput(input)
	world4D := World4DFromInput(input)

	var part1, part2 int
	for i := 0; i < 6; i++ {
		part1 = world3D.Step()
		part2 = world4D.Step()
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
