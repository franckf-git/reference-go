package main

import (
	"strconv"
	"strings"
)

// A sequence is a unique sequence of rule indexes to match
type Sequence []int

// A rule can is composed of one or several sequences
type Rule struct {
	Terminal  byte
	Sequences []Sequence
}

type Grammar []Rule

func parseRules(specs []string) Grammar {
	rules := make([]Rule, 200)
	for _, spec := range specs {
		s := strings.Split(spec, ":")
		if len(s) != 2 {
			panic("invalid input")
		}
		index, _ := strconv.Atoi(s[0])
		if i := strings.IndexAny(s[1], "ab"); i != -1 {
			char := s[1][i]
			rules[index] = Rule{Terminal: char}
		} else {
			rule := Rule{}
			seqs := strings.Split(s[1], "|")
			rule.Sequences = make([]Sequence, len(seqs))
			for i, seqspec := range seqs {
				fields := strings.Fields(seqspec)
				rule.Sequences[i] = make(Sequence, len(fields))
				for f, field := range fields {
					var err error
					rule.Sequences[i][f], err = strconv.Atoi(field)
					if err != nil {
						panic(err)
					}
				}
			}
			rules[index] = rule
		}
	}
	return rules
}

func parseInput(input []string) ([]Rule, []string) {
	var pivot int
	// Find the split point
	for i, line := range input {
		if line == "" {
			pivot = i
			break
		}
	}
	return parseRules(input[:pivot]), input[pivot+1:]
}
