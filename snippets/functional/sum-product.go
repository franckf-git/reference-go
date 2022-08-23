package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Reduce[T any](
	slice []T,
	initialValue T,
	callback func(previousValue T, currentValue T, currentIndex int) T,
) T

func Sum[T constraints.Integer | constraints.Float](slice []T) T {
	callback := func(previousValue T, currentValue T, currentIndex int) T {
		return previousValue + currentValue
	}

	return Reduce(slice, 0, callback)
}

func Product[T constraints.Integer | constraints.Float](slice []T) T {
	callback := func(previousValue T, currentValue T, currentIndex int) T {
		return previousValue * currentValue
	}

	return Reduce(slice, 1, callback)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf(
		"Sum: %d\nProduct: %d\n",
		Sum(numbers),
		Product(numbers),
	)
}
