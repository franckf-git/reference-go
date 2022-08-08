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
	arrRoman := strings.Split(roman, "")
	for i := 0; i < len(arrRoman); i++ {
		// si avant dernier I
		if arrRoman[i] == "I" && i+1 == len(arrRoman)-1 {
			// et si dernier V > donc IV - on finit sur plus 4
			if arrRoman[i+1] == "V" {
				number += 4
				break
			}
			// et si dernier X > donc IX - on finit sur plus 9
			if arrRoman[i+1] == "X" {
				number += 9
				break
			}
		}
		// si avant dernier X - Start to be ugly
		if arrRoman[i] == "X" && i+1 <= len(arrRoman)-1 {
			// et si dernier L > donc XL - on finit sur plus 40
			if arrRoman[i+1] == "L" {
				number += 40
				break
			}
		}
		if arrRoman[i] == "V" {
			number += 4
		}
		if arrRoman[i] == "X" {
			number += 9
		}
		if arrRoman[i] == "L" {
			number += 49
		}
		number++
	}

	/*
		var arrRomanRev = make([]string, 0)
		for i := len(arrRoman) - 1; i >= 0; i-- {
			arrRomanRev = append(arrRomanRev, arrRoman[i])
		}

		for i, v := range arrRomanRev {
			switch {
			case v == "V" && arrRomanRev[i+1] == "I":
				number += 4
			case v == "V":
				number += 5
			}
			if v == "I" {
				number++
			}
		}
	*/
	return
}
