package main

import "fmt"

func main() {
	pyramidByHigh(51)
}

func pyramidByHigh(steps int) {
	pyramidSteps := make([]string, steps)
	for index, step := range pyramidSteps {
		var dot string
		size := index * 2
		for blank := 0; blank < steps-index; blank++ {
			dot = dot + " "

		}
		for character := 0; character <= size; character++ {
			dot = dot + "."
		}

		step = dot
		fmt.Println(step)
	}

}
