package main

import "fmt"

// ProcessNum transform number into collatz
func ProcessNum(number int) int {
	if number%2 == 0 {
		return number / 2
	}
	return (number * 3) + 1

}

// CollatzProcess process a number until nothing is left
func CollatzProcess(number int) []int {
	result := []int{}
	entryPoint := []int{number}
	result = append(result, entryPoint...)
	for result[len(result)-1] > 1 {
		numberProcessed := ProcessNum(result[len(result)-1])
		newEntry := []int{numberProcessed}
		result = append(result, newEntry...)
	}
	return result
}

func main() {
	result := CollatzProcess(17)
	fmt.Print(result)
}
