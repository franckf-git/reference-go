package main

import "fmt"

// Fizzbuzz - you know it
func Fizzbuzz(number int) string {
	var output string
	for i := 1; i <= number; i++ {
		if i%5 == 0 && i%3 == 0 {
			output += "FizzBuzz"
		} else if i%3 == 0 {
			output += "Fizz"
		} else if i%5 == 0 {
			output += "Buzz"
		} else {
			output += fmt.Sprint(i)
		}
		output += " "
	}
	return output
}

func main() {
	result := Fizzbuzz(16)
	fmt.Print(result)
}
