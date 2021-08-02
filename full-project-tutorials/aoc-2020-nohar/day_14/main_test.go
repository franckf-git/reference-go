package main

import (
	"reflect"
	"testing"

	"gitlab.com/neuware/aoc-2020/utils"
)

func TestMaskApply(t *testing.T) {
	cases := []struct {
		Mask     string
		Input    uint64
		Expected uint64
	}{
		{
			Mask:     "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			Input:    11,
			Expected: 73,
		},
		{
			Mask:     "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			Input:    101,
			Expected: 101,
		},
		{
			Mask:     "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			Input:    0,
			Expected: 64,
		},
	}

	for _, c := range cases {
		m := ParseMask(c.Mask)
		res := m.Apply(c.Input)
		if res != c.Expected {
			t.Errorf(
				"Mask %s, Input %b, Expected %b, Got %b",
				c.Mask, c.Input, c.Expected, res,
			)
		}
	}
}

func TestMaskGenerate(t *testing.T) {
	cases := []struct {
		Mask     string
		Input    uint64
		Expected []uint64
	}{
		{
			Mask:     "000000000000000000000000000000X1001X",
			Input:    42,
			Expected: []uint64{26, 27, 58, 59},
		},
		{
			Mask:     "00000000000000000000000000000000X0XX",
			Input:    26,
			Expected: []uint64{16, 17, 18, 19, 24, 25, 26, 27},
		},
	}

	for _, c := range cases {
		m := ParseMask(c.Mask)
		res := m.Generate(c.Input)
		if !reflect.DeepEqual(c.Expected, res) {
			t.Errorf(
				"Mask %s, Input %b, Expected %v, Got %v",
				c.Mask, c.Input, c.Expected, res,
			)
		}
	}
}

func TestPart1(t *testing.T) {
	cases := []struct {
		Program  []string
		Expected uint64
	}{
		{
			Program: []string{
				"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
				"mem[8] = 11",
				"mem[7] = 101",
				"mem[8] = 0",
			},
			Expected: 165,
		},
	}

	for i, c := range cases {
		res := run(c.Program, part1)
		if res != c.Expected {
			t.Errorf("Case #%d, Expected %d, Got %d", i+1, c.Expected, res)
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		Program  []string
		Expected uint64
	}{
		{
			Program: []string{
				"mask = 000000000000000000000000000000X1001X",
				"mem[42] = 100",
				"mask = 00000000000000000000000000000000X0XX",
				"mem[26] = 1",
			},
			Expected: 208,
		},
	}

	for i, c := range cases {
		res := run(c.Program, part2)
		if res != c.Expected {
			t.Errorf("Case #%d, Expected %d, Got %d", i+1, c.Expected, res)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	program := utils.ReadLines()
	for n := 0; n < b.N; n++ {
		if run(program, part1) == 0 {
			b.Fail()
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	program := utils.ReadLines()
	for n := 0; n < b.N; n++ {
		if run(program, part2) == 0 {
			b.Fail()
		}
	}
}
