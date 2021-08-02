package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gitlab.com/neuware/aoc-2020/utils"
)

type Passport map[string]string
type Validator func(string) bool

var (
	validators = map[string]Validator{
		"byr": yearValidator(1920, 2002),
		"iyr": yearValidator(2010, 2020),
		"eyr": yearValidator(2020, 2030),
		"hgt": isValidHeight,
		"hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString,
		"ecl": regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`).MatchString,
		"pid": regexp.MustCompile(`^\d{9}$`).MatchString,
	}
	heightFmt = regexp.MustCompile(`^(\d+)(cm|in)$`)
)

func main() {
	var part1, part2 int

	iterPassports(utils.ReadLines(), func(p Passport) {
		if isValidPart1(p) {
			part1++
		}
		if isValidPart2(p) {
			part2++
		}
	})

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func isValidPart1(p Passport) bool {
	for key, _ := range validators {
		if _, ok := p[key]; !ok {
			return false
		}
	}
	return true
}

func isValidPart2(p Passport) bool {
	for key, valid := range validators {
		if s, ok := p[key]; !ok || !valid(s) {
			return false
		}
	}
	return true
}

func iterPassports(input []string, cbk func(Passport)) {
	p := make(Passport)
	for _, line := range input {
		// Parse the passport
		for _, field := range strings.Fields(line) {
			kv := strings.SplitN(field, ":", 2)
			if len(kv) == 2 {
				p[kv[0]] = kv[1]
			}
		}

		// On empty lines, process the passport and reset it
		if line == "" {
			cbk(p)
			p = make(Passport)
		}
	}
	// Process the remaining passport
	cbk(p)
}

func yearValidator(min, max int) Validator {
	return func(s string) bool {
		y, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return min <= y && y <= max
	}
}

func isValidHeight(s string) bool {
	m := heightFmt.FindStringSubmatch(s)
	if m == nil {
		return false
	}
	h, _ := strconv.Atoi(m[1])
	switch m[2] {
	case "cm":
		return 150 <= h && h <= 193
	case "in":
		return 59 <= h && h <= 76
	}
	return false
}
