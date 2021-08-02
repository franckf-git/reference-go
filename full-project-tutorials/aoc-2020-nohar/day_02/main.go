package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"gitlab.com/neuware/aoc-2020/utils"
)

var entryFmt = regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)

type PasswordPolicy func(string, byte, int, int) bool

func isValid(entry string, f PasswordPolicy) bool {
	res := entryFmt.FindStringSubmatch(entry)
	if res == nil {
		fmt.Fprintln(os.Stderr, "Malformed input: ", entry)
		return false
	}
	i1, _ := strconv.Atoi(res[1])
	i2, _ := strconv.Atoi(res[2])
	char := res[3][0]
	password := res[4]
	return f(password, char, i1, i2)
}

func partOne(password string, char byte, min, max int) bool {
	var count int
	for _, c := range []byte(password) {
		if char == c {
			count++
		}
	}
	return min <= count && count <= max
}

func partTwo(password string, char byte, p1, p2 int) bool {
	p1--
	p2--
	if 0 > p1 || p1 >= len(password) || 0 > p2 || p2 >= len(password) {
		return false
	}
	return (password[p1] == char) != (password[p2] == char)
}

func main() {
	var count, count2 int
	for _, line := range utils.ReadLines() {
		if isValid(line, partOne) {
			count++
		}
		if isValid(line, partTwo) {
			count2++
		}
	}
	fmt.Println("Part1:", count)
	fmt.Println("Part2:", count2)
}
