package katas

import (
	"strings"
)

type RomanCharacter struct {
	decimal int
	roman   string
}

// must be in decr order
var RomanCharacters = []RomanCharacter{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func num2rom(number int) (roman string) {
	var result strings.Builder

	for _, v := range RomanCharacters {
		for number >= v.decimal {
			result.WriteString(v.roman)
			number -= v.decimal
		}
	}

	roman = result.String()
	return
}

func rom2num(roman string) (number int) {
	number = 1
	return
}
