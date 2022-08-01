package katas

import (
	"fmt"
	"strings"
)

func num2rom(number int) (roman string) {
	var result strings.Builder

	// X
	if number >= 10 {
		result.WriteString("X")
		number = number - 10
	}
	if fmt.Sprint(number)[len(fmt.Sprint(number))-1:] == "9" {
		result.WriteString("IX")
		number = number - 9
	}
	// V
	if number >= 5 {
		result.WriteString("V")
		number = number - 5
	}
	if fmt.Sprint(number)[len(fmt.Sprint(number))-1:] == "4" {
		result.WriteString("IV")
		number = number - 4
		print(number)
	}

	for i := 0; i < number; i++ {
		result.WriteString("I")
	}

	roman = result.String()
	return
}

func rom2num(roman string) (number int) {
	number = 1
	return
}
