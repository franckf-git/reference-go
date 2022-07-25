package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("--")
	fmt.Print(printTriforce(5))
	fmt.Println("--")
}

var symbol string = "▲"

func printTriforce(base int) (result string) {
	result = "\n"
	if base%2 == 0 {
		result = "cannot use even number - only odd"
		return
	}
	part1 := make([]string, base/2+1)
	indent := len(part1) - 1
	for i := range part1 {
		line := i + 1
		currentSymbol := strings.Repeat(symbol, line*2-1)
		part1[i] = part1[i] + strings.Repeat(" ", indent)
		part1[i] = part1[i] + currentSymbol
		indent--
	}
	// ---
	for _, v := range part1 {
		result = result + strings.Repeat(" ", len(part1)) + v + "\n"
	}
	invIndent := len(part1)
	for _, v := range part1 {
		result = result + v + strings.Repeat(" ", invIndent) + v + "\n"
		invIndent--
	}
	return
}
