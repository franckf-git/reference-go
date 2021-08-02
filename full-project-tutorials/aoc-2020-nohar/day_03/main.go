package main

import (
	"fmt"

	"gitlab.com/neuware/aoc-2020/utils"
)

func countTrees(input []string, right, down int) int {
	height, width := len(input), len(input[0])
	var count, x int
	for y := 0; y < height; y += down {
		if input[y][x] == '#' {
			count++
		}
		x = (x + right) % width
	}
	return count
}

func main() {
	input := utils.ReadLines()

	result := countTrees(input, 3, 1)
	fmt.Println("Part 1:", result)

	result *= countTrees(input, 1, 1)
	result *= countTrees(input, 5, 1)
	result *= countTrees(input, 7, 1)
	result *= countTrees(input, 1, 2)
	fmt.Println("Part 2:", result)
}
