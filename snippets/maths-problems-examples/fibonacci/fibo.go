package main

// Fibo return fibonacci suite on number th
func Fibo(number int) []int {
	result := []int{1, 1}
	for len(result) < number {
		nMinus1 := result[len(result)-1]
		nMinus2 := result[len(result)-2]
		current := nMinus1 + nMinus2
		newEntry := []int{current}
		result = append(result, newEntry...)
	}
	return result
}
