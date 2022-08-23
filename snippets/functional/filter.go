package main

import "fmt"

func Filter[T any](
	slice []T,
	callback func(value T, index int) bool,
) []T {
	filteredSlice := make([]T, 0, len(slice))

	for i, v := range slice {
		if callback(v, i) {
			filteredSlice = append(filteredSlice, v)
		}
	}

	return filteredSlice
}

func main() {
	numbers := []float32{0.12, 0.97, 0.26, 0.83, 0.43, 0.04, 0.92}

	filteredNumbers := Filter(numbers, func(value float32, index int) bool {
		return value > 0.5
	})

	fmt.Println(filteredNumbers)
	fmt.Printf("Filtered numbers count: %d\n", len(filteredNumbers))
	fmt.Printf("Removed numbers count: %d\n", len(numbers)-len(filteredNumbers))
}