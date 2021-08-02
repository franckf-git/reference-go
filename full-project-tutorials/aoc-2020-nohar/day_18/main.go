package main

import (
	"fmt"

	"gitlab.com/neuware/aoc-2020/utils"
)

type Binop func(int, int) int

func add(a, b int) int { return a + b }
func mul(a, b int) int { return a * b }

func evalAt(input string, index int) (result, pos int) {
	var op Binop
	for pos = index; pos < len(input); pos++ {
		switch c := input[pos]; c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if op != nil {
				result = op(result, int(c-'0'))
			} else {
				result = int(c - '0')
			}
		case '(':
			var x int
			x, pos = evalAt(input, pos+1)
			if op != nil {
				result = op(result, x)
			} else {
				result = x
			}
		case ')':
			return result, pos
		case '+':
			op = add
		case '*':
			op = mul
		default:
		}
	}
	return
}

func evalPart1(expr string) int {
	result, _ := evalAt(expr, 0)
	return result
}

func evalSum(input string, index int) (result, pos int) {
	for pos = index; pos < len(input); pos++ {
		switch c := input[pos]; c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			result += int(c - '0')
		case '(':
			var x int
			x, pos = evalProduct(input, pos+1)
			result += x
		case ')', '*':
			return result, pos - 1
		default:
		}
	}
	return
}

func evalProduct(input string, index int) (result, pos int) {
	result = 1
	for pos = index; pos < len(input); pos++ {
		switch c := input[pos]; c {
		case '(', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			var x int
			x, pos = evalSum(input, pos)
			result *= x
		case ')':
			return result, pos
		default:
		}
	}
	return
}

func evalPart2(input string) int {
	result, _ := evalProduct(input, 0)
	return result
}

func main() {
	input := utils.ReadLines()

	var part1, part2 int
	for _, expr := range input {
		part1 += evalPart1(expr)
		part2 += evalPart2(expr)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
