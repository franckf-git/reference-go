package main

import (
	"fmt"
	"math/bits"
	"strings"

	"gitlab.com/neuware/aoc-2020/utils"
)

type Ticket []int
type Range struct{ Min, Max int }
type Field struct {
	Name string
	Low  Range
	High Range
	Mask uint32
}

func (r Range) Contains(value int) bool {
	return r.Min <= value && value <= r.Max
}

func (f Field) InRange(value int) bool {
	return f.Low.Contains(value) || f.High.Contains(value)
}

func wantedFields(fields []Field) (mask uint32) {
	for _, r := range fields {
		if strings.HasPrefix(r.Name, "departure") {
			mask |= r.Mask
		}
	}
	return
}

func findCandidates(fields []Field, tickets []Ticket) []uint32 {
	candidates := make([]uint32, len(fields))
	for i := range candidates {
	FieldsLoop:
		for _, field := range fields {
			for _, ticket := range tickets {
				if !field.InRange(ticket[i]) {
					continue FieldsLoop
				}
			}
			candidates[i] |= field.Mask
		}
	}
	return candidates
}

func part2(fields []Field, ticket Ticket, nearby []Ticket) int {
	candidates := findCandidates(fields, nearby)
	wanted := wantedFields(fields)

	// Refine results until all the fields we want are unambiguously known.
	result := 1
	var known uint32
	for known&wanted != wanted {
		found := known
		for i, c := range candidates {
			if c == 0 {
				// We've already found this field, skip it
				continue
			}

			// Prune all known fields from this set of candidates
			c &^= found

			// There's only one possibility left
			if bits.OnesCount32(c) == 1 {
				if c&wanted != 0 {
					result *= ticket[i]
				}
				found |= c
				candidates[i] = 0 // Mark as found
			}
		}
		if found == known {
			panic("couldn't find solution")
		}
		known = found
	}
	return result
}

func isValid(ticket Ticket, fields []Field) (ok bool, rate int) {
	ok = true
OuterLoop:
	for _, v := range ticket {
		for _, r := range fields {
			if r.InRange(v) {
				continue OuterLoop
			}
		}
		ok = false
		rate += v
	}
	return
}

func validateTickets(tickets []Ticket, fields []Field) ([]Ticket, int) {
	var sum int
	valid := tickets[:0]
	for _, ticket := range tickets {
		if ok, rate := isValid(ticket, fields); !ok {
			sum += rate
			continue
		}
		valid = append(valid, ticket)
	}
	return valid, sum
}

func main() {
	fields, ticket, nearby := parseInput(utils.ReadLines())
	nearby, part1 := validateTickets(nearby, fields)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2(fields, ticket, nearby))
}
