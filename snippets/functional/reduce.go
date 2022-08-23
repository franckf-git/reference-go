package main

import "fmt"

func Reduce[T any](
	slice []T,
	initialValue T,
	callback func(previousValue T, currentValue T, currentIndex int) T,
) T {
	reducedValue := initialValue

	for i, v := range slice {
		reducedValue = callback(reducedValue, v, i)
	}

	return reducedValue
}

func main() {
	numbers := []float32{0.12, 0.97, 0.26, 0.83, 0.43, 0.04, 0.92}

	sumOfNumbers := Reduce(numbers, 0, func(previousValue, currentValue float32, index int) float32 {
		return previousValue + currentValue
	})

	fmt.Printf("Reduced value: %.2f\n", sumOfNumbers)
}