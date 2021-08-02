package main

import "testing"

func TestWorld3D(t *testing.T) {
	world := World3DFromInput([]string{".#.", "..#", "###"})
	if len(world) != 5 {
		t.Fatalf("initial step, expected 5 active cells, got %d", len(world))
	}

	for i, expected := range []int{11, 21, 38} {
		res := world.Step()
		if res != expected {
			t.Fatalf(
				"cycle %d, expected %d active cells, got %d (%v)",
				i+1, expected, res, world,
			)
		}
	}
}
