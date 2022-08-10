package main

import "testing"

func Test_tac(t *testing.T) {
	in := []byte(`line1
line2
line3`)

	out := `line3
line2
line1`
	want := tac(in)
	if string(want) != out {
		t.Error("different output, have: ", string(want))
	}
}
