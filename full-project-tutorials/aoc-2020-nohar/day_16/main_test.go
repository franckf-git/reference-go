package main

import (
	"testing"

	"gitlab.com/neuware/aoc-2020/utils"
)

var example = []string{
	"class: 1-3 or 5-7",
	"row: 6-11 or 33-44",
	"seat: 13-40 or 45-50",
	"",
	"your ticket:",
	"7,1,14",
	"",
	"nearby tickets:",
	"7,3,47",
	"40,4,50",
	"55,2,20",
	"38,6,12",
}

func TestPart1(t *testing.T) {
	fields, _, nearby := parseInput(example)
	_, res := validateTickets(nearby, fields)
	if res != 71 {
		t.Fatalf("Expected 71, Got %d", res)
	}
}

func TestValidate(t *testing.T) {
	fields, _, nearby := parseInput(example)
	valid, res := validateTickets(nearby, fields)
	if len(valid) == len(nearby) {
		t.Fatal("No invalid tickets")
	}
	valid, res = validateTickets(valid, fields)
	if res != 0 {
		t.Fatalf("expected res=0, got %d", res)
	}
}

func BenchmarkPart2(b *testing.B) {
	fields, ticket, nearby := parseInput(utils.ReadLines())
	nearby, _ = validateTickets(nearby, fields)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if part2(fields, ticket, nearby) == 0 {
			b.Fail()
		}
	}
}
