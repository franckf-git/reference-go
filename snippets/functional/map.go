package main

import "fmt"

func Map[T any](
	slice []T,
	callback func(value T, index int) T,
) []T {
	mappedSlice := make([]T, len(slice))

	for i, v := range slice {
		mappedSlice[i] = callback(v, i)
	}

	return mappedSlice
}

func main() {
	x := []string{"top", "log", "pend"}

	y := Map(x, func(value string, index int) string {
		return "s" + value
	})

	fmt.Println(y)
}