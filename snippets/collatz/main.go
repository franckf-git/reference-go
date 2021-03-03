package main

import "fmt"

// ProcessNum transform number into collatz
func ProcessNum(number int) int {
	return number
}

// CollatzProcess process a number until nothing is left
func CollatzProcess(number int) []int {
	res := []int{1, 2, 3}
	return res
}

func main() {
	result := CollatzProcess(16)
	fmt.Print(result)
}
