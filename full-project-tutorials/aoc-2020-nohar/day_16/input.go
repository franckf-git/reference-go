package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	rangeFmt         = regexp.MustCompile(`^(\d+)-(\d+)$`)
	errInvalidField  = errors.New("invalid rule")
	errInvalidTicket = errors.New("invalid ticket")
)

func parseRange(input string) Range {
	groups := rangeFmt.FindStringSubmatch(input)
	if groups == nil {
		panic(errInvalidField)
	}
	min, _ := strconv.Atoi(groups[1])
	max, _ := strconv.Atoi(groups[2])
	return Range{min, max}
}

func parseField(spec string) Field {
	s := strings.SplitN(spec, ": ", 2)
	if len(s) != 2 {
		panic(errInvalidField)
	}
	r := Field{Name: s[0]}
	ranges := strings.Split(s[1], " or ")
	if len(ranges) != 2 {
		panic(errInvalidField)
	}
	r.Low = parseRange(ranges[0])
	r.High = parseRange(ranges[1])
	return r
}

func parseTicket(line string) Ticket {
	s := strings.Split(line, ",")
	t := make(Ticket, len(s))
	for i, field := range s {
		var err error
		if t[i], err = strconv.Atoi(field); err != nil {
			panic(errInvalidTicket)
		}
	}
	return t
}

func parseInput(input []string) ([]Field, Ticket, []Ticket) {
	// Parse validation rules
	rules := make([]Field, 0, 20)
	for i, line := range input {
		if line == "" {
			// Skip the "your ticket:" line
			input = input[i+2:]
			break
		}
		r := parseField(line)
		r.Mask = 1 << i
		rules = append(rules, r)
	}

	// Parse "our" ticket
	ticket := parseTicket(input[0])
	input = input[3:]

	// Parse "nearby" tickets
	nearby := make([]Ticket, len(input))
	for i, line := range input {
		nearby[i] = parseTicket(line)
	}

	return rules, ticket, nearby
}
