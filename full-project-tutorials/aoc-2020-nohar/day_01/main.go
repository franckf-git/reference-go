package main

import (
	"fmt"

	"gitlab.com/neuware/aoc-2020/utils"
)

func partOne(input []int) {
	for i, x := range input {
		for j, y := range input {
			if i != j && x+y == 2020 {
				fmt.Println("Part 1: ", x*y)
				return
			}
		}
	}
}

func partTwo(input []int) {
	for i, x := range input {
		for j, y := range input {
			if i == j {
				continue
			}
			for k, z := range input {
				if j != k && x+y+z == 2020 {
					fmt.Println("Part 2: ", x*y*z)
					return
				}
			}
		}
	}
}

func main() {
	input := utils.ReadInts()
	partOne(input)
	partTwo(input)
}
