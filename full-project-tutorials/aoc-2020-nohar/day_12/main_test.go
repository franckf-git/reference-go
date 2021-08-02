package main

import "testing"

func TestNavigatePart1(t *testing.T) {
	cases := []struct {
		Input []string
		ResX  int
		ResY  int
	}{
		{
			Input: []string{"F10", "N3", "F7", "R90", "F11"},
			ResX:  17,
			ResY:  -8,
		},
	}
	for i, c := range cases {
		x, y := navigatePart1(c.Input)
		if x != c.ResX || y != c.ResY {
			t.Errorf(
				"Case %d: expected (%d, %d), got (%d, %d)",
				i+1, c.ResX, c.ResY, x, y,
			)
		}
	}
}

func TestNavigatePart2(t *testing.T) {
	cases := []struct {
		Input []string
		ResX  int
		ResY  int
	}{
		{
			Input: []string{"F10", "N3", "F7", "R90", "F11"},
			ResX:  214,
			ResY:  -72,
		},
	}
	for i, c := range cases {
		x, y := navigatePart2(c.Input)
		if x != c.ResX || y != c.ResY {
			t.Errorf(
				"Case %d: expected (%d, %d), got (%d, %d)",
				i+1, c.ResX, c.ResY, x, y,
			)
		}
	}
}
