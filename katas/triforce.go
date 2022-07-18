package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(printTriforce(3))
}

var symbol string = "▲"

func printTriforce(base int) (result string) {
	if base%2 == 0 {
		result = "cannot use even number - only odd"
		return
	}
	indent := base
	lines := base + 1
	for i := 0; i < lines; i++ {
		currentSymbol := symbol
		if i%2 != 0 {
			currentSymbol = strings.Repeat(symbol, 3)
		}
		currentLine := fmt.Sprintln(strings.Repeat(" ", indent), currentSymbol)
		result = result + currentLine
		indent--
	}
	return
}
