package main

import (
	"testing"

	"gitlab.com/neuware/aoc-2020/utils"
)

func TestMatchPart1(t *testing.T) {
	input := []string{
		"0: 4 1 5",
		"1: 2 3 | 3 2",
		"2: 4 4 | 5 5",
		"3: 4 5 | 5 4",
		`4: "a"`,
		`5: "b"`,
		"",
		"ababbb",
		"bababa",
		"abbbab",
		"aaabbb",
		"aaaabbb",
	}

	rules, input := parseInput(input)
	res := countMatches(input, rules)
	if res != 2 {
		t.Fatalf("Expected %d, Got %d", 2, res)
	}
}

func TestMatchPart2(t *testing.T) {
	input := []string{
		`42: 9 14 | 10 1`,
		`9: 14 27 | 1 26`,
		`10: 23 14 | 28 1`,
		`1: "a"`,
		`11: 42 31`,
		`5: 1 14 | 15 1`,
		`19: 14 1 | 14 14`,
		`12: 24 14 | 19 1`,
		`16: 15 1 | 14 14`,
		`31: 14 17 | 1 13`,
		`6: 14 14 | 1 14`,
		`2: 1 24 | 14 4`,
		`0: 8 11`,
		`13: 14 3 | 1 12`,
		`15: 1 | 14`,
		`17: 14 2 | 1 7`,
		`23: 25 1 | 22 14`,
		`28: 16 1`,
		`4: 1 1`,
		`20: 14 14 | 1 15`,
		`3: 5 14 | 16 1`,
		`27: 1 6 | 14 18`,
		`14: "b"`,
		`21: 14 1 | 1 14`,
		`25: 1 1 | 1 14`,
		`22: 14 14`,
		`8: 42`,
		`26: 14 22 | 1 20`,
		`18: 15 15`,
		`7: 14 5 | 1 21`,
		`24: 14 1`,
		``,
		`abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa`,
		`bbabbbbaabaabba`,
		`babbbbaabbbbbabbbbbbaabaaabaaa`,
		`aaabbbbbbaaaabaababaabababbabaaabbababababaaa`,
		`bbbbbbbaaaabbbbaaabbabaaa`,
		`bbbababbbbaaaaaaaabbababaaababaabab`,
		`ababaaaaaabaaab`,
		`ababaaaaabbbaba`,
		`baabbaaaabbaaaababbaababb`,
		`abbbbabbbbaaaababbbbbbaaaababb`,
		`aaaaabbaabaaaaababaa`,
		`aaaabbaaaabbaaa`,
		`aaaabbaabbaaaaaaabbbabbbaaabbaabaaa`,
		`babaaabbbaaabaababbaabababaaab`,
		`aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`,
	}

	rules, input := parseInput(input)
	rules = patch(rules)
	res := countMatches(input, rules)
	if res != 12 {
		t.Fatalf("Expected %d, Got %d", 12, res)
	}
}

func BenchmarkPart1(b *testing.B) {
	rules, input := parseInput(utils.ReadLines())
	var size int = len(input)
	var res int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if match(input[n%size], rules) {
			res++
		}
	}
	if res == 0 {
		b.Fail()
	}
}

func BenchmarkPart2(b *testing.B) {
	rules, input := parseInput(utils.ReadLines())
	rules = patch(rules)
	var size int = len(input)
	var res int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if match(input[n%size], rules) {
			res++
		}
	}
	if res == 0 {
		b.Fail()
	}
}
