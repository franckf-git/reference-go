package main

import (
	"reflect"
	"testing"

	"gitlab.com/neuware/aoc-2020/utils"
)

func TestParseSerialize(t *testing.T) {
	sampleInput := []string{
		"L.LL.LL.LL",
		"LL####L.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.#L.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LL##LL.L",
		"L.LLLLL.LL",
	}
	grid := NewGridFromInput(sampleInput)
	res := grid.ToSlice()
	if !reflect.DeepEqual(res, sampleInput) {
		t.Fatalf("Expected: %v\nGot: %v", sampleInput, res)
	}
}

func TestStepPart1(t *testing.T) {
	type caseStep struct {
		Result  []string
		Changes int
	}
	cases := []struct {
		Init  []string
		Steps []caseStep
	}{
		{
			Init: []string{"LLL", "LLL", "LLL"},
			Steps: []caseStep{
				{
					Result:  []string{"###", "###", "###"},
					Changes: 9,
				},
				{
					Result:  []string{"#L#", "LLL", "#L#"},
					Changes: 5,
				},
				{
					Result:  []string{"#L#", "LLL", "#L#"},
					Changes: 0,
				},
			},
		},
	}

	for i, c := range cases {
		src := NewGridFromInput(c.Init)
		dst := NewGrid(src.Width(), src.Height())
		var changes int
		for j, exp := range c.Steps {
			changes = step(src, dst, CanTakeSeatPart1, MustLeaveSeatPart1)
			if changes != exp.Changes {
				t.Errorf("case %d, step %d: expected %d changes, got %d", i, j, exp.Changes, changes)
			}
			result := dst.ToSlice()
			if !reflect.DeepEqual(exp.Result, result) {
				t.Fatalf("case %d, step %d:\nExpected:\n%v\nGot:\n%v", i, j, exp.Result, result)
			}
			src, dst = dst, src
		}
	}
}

func TestStepPart2(t *testing.T) {
	type caseStep struct {
		Result  []string
		Changes int
	}
	cases := []struct {
		Init  []string
		Steps []caseStep
	}{
		{
			Init: []string{
				"L.LL.LL.LL",
				"LLLLLLL.LL",
				"L.L.L..L..",
				"LLLL.LL.LL",
				"L.LL.LL.LL",
				"L.LLLLL.LL",
				"..L.L.....",
				"LLLLLLLLLL",
				"L.LLLLLL.L",
				"L.LLLLL.LL",
			},
			Steps: []caseStep{
				{
					Result: []string{
						"#.##.##.##", "#######.##", "#.#.#..#..", "####.##.##", "#.##.##.##",
						"#.#####.##", "..#.#.....", "##########", "#.######.#", "#.#####.##",
					},
					Changes: 71,
				},
				{
					Result: []string{
						"#.LL.LL.L#", "#LLLLLL.LL", "L.L.L..L..", "LLLL.LL.LL", "L.LL.LL.LL",
						"L.LLLLL.LL", "..L.L.....", "LLLLLLLLL#", "#.LLLLLL.L", "#.LLLLL.L#",
					},
					Changes: 64,
				},
				{
					Result: []string{
						"#.L#.##.L#", "#L#####.LL", "L.#.#..#..", "##L#.##.##", "#.##.#L.##",
						"#.#####.#L", "..#.#.....", "LLL####LL#", "#.L#####.L", "#.L####.L#",
					},
					Changes: 46,
				},
				{
					Result: []string{
						"#.L#.L#.L#", "#LLLLLL.LL", "L.L.L..#..", "##LL.LL.L#", "L.LL.LL.L#",
						"#.LLLLL.LL", "..L.L.....", "LLLLLLLLL#", "#.LLLLL#.L", "#.L#LL#.L#",
					},
					Changes: 35,
				},
				{
					Result: []string{
						"#.L#.L#.L#", "#LLLLLL.LL", "L.L.L..#..", "##L#.#L.L#", "L.L#.#L.L#",
						"#.L####.LL", "..#.#.....", "LLL###LLL#", "#.LLLLL#.L", "#.L#LL#.L#",
					},
					Changes: 13,
				},
				{
					Result: []string{
						"#.L#.L#.L#", "#LLLLLL.LL", "L.L.L..#..", "##L#.#L.L#", "L.L#.LL.L#",
						"#.LLLL#.LL", "..#.L.....", "LLL###LLL#", "#.LLLLL#.L", "#.L#LL#.L#",
					},
					Changes: 5,
				},
				{
					Result: []string{
						"#.L#.L#.L#", "#LLLLLL.LL", "L.L.L..#..", "##L#.#L.L#", "L.L#.LL.L#",
						"#.LLLL#.LL", "..#.L.....", "LLL###LLL#", "#.LLLLL#.L", "#.L#LL#.L#",
					},
					Changes: 0,
				},
			},
		},
	}

	for i, c := range cases {
		src := NewGridFromInput(c.Init)
		dst := NewGrid(src.Width(), src.Height())
		var changes int
		for j, exp := range c.Steps {
			changes = step(src, dst, CanTakeSeatPart2, MustLeaveSeatPart2)
			if changes != exp.Changes {
				t.Errorf("case %d, step %d: expected %d changes, got %d", i, j+1, exp.Changes, changes)
			}
			result := dst.ToSlice()
			if !reflect.DeepEqual(exp.Result, result) {
				t.Fatalf("case %d, step %d:\nExpected:\n%v\nGot:\n%v", i, j+1, exp.Result, result)
			}
			src, dst = dst, src
		}
	}
}

func BenchmarkStepPart1(b *testing.B) {
	src := NewGridFromInput(utils.ReadLines())
	dst := NewGrid(src.Width(), src.Height())
	for i := 0; i < 3; i++ {
		step(src, dst, CanTakeSeatPart1, MustLeaveSeatPart1)
		src, dst = dst, src
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		changes := step(src, dst, CanTakeSeatPart1, MustLeaveSeatPart1)
		if changes == 0 {
			b.Fail()
		}
	}
}

func BenchmarkStepPart2(b *testing.B) {
	src := NewGridFromInput(utils.ReadLines())
	dst := NewGrid(src.Width(), src.Height())
	for i := 0; i < 3; i++ {
		step(src, dst, CanTakeSeatPart2, MustLeaveSeatPart2)
		src, dst = dst, src
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		changes := step(src, dst, CanTakeSeatPart2, MustLeaveSeatPart2)
		if changes == 0 {
			b.Fail()
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	input := utils.ReadLines()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		g, _ := run(input, CanTakeSeatPart1, MustLeaveSeatPart1)
		if g == nil {
			b.Fail()
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	input := utils.ReadLines()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		g, _ := run(input, CanTakeSeatPart2, MustLeaveSeatPart2)
		if g == nil {
			b.Fail()
		}
	}
}
