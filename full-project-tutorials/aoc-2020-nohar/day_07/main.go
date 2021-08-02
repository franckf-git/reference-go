package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gitlab.com/neuware/aoc-2020/utils"
)

// Color type of a bag
type Color string

// Rule models how many bags of each type we can contain
type Rule struct {
	Number int
	Color  Color
}

// RuleGraph models the dependency tree on bag colors.
type RuleGraph map[Color][]Rule

// HasAncestor returns true if a bag with color 'c' can be (transitively)
// contained in a bag with color 'a'.
func (g RuleGraph) HasAncestor(c, a Color) bool {
	rules := g[a]
	for _, rule := range rules {
		if rule.Color == c {
			return true
		}
	}
	for _, rule := range rules {
		if g.HasAncestor(c, rule.Color) {
			return true
		}
	}
	return false
}

// CountInner returns the total number of bags that must be contained in a
// bag with color 'key'.
func (g RuleGraph) CountInner(key Color) int {
	var count int
	for _, rule := range g[key] {
		count += rule.Number * (1 + g.CountInner(rule.Color))
	}
	return count
}

func main() {
	var part1 int
	graph := parseRules(utils.ReadLines())
	c := Color("shiny gold")
	for key := range graph {
		if graph.HasAncestor(c, key) {
			part1++
		}
	}
	part2 := graph.CountInner(c)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

var bagRuleFmt = regexp.MustCompile(`(\d) ([a-z]+ [a-z]+) bag`)

func parseRules(input []string) RuleGraph {
	g := make(RuleGraph)
	for _, line := range utils.ReadLines() {
		ls := strings.SplitN(line, " bags contain ", 2)
		if len(ls) != 2 {
			utils.Fatal("invalid rule:", line)
		}
		key := Color(ls[0])
		matches := bagRuleFmt.FindAllStringSubmatch(ls[1], -1)
		if matches == nil {
			continue
		}
		g[key] = make([]Rule, len(matches))
		for i, m := range matches {
			b := Rule{}
			b.Number, _ = strconv.Atoi(m[1])
			b.Color = Color(m[2])
			g[key][i] = b
		}
	}
	return g
}
