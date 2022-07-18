package main

import "testing"

func TestTriForce3(t *testing.T) {
	expected := `
   ▲
  ▲▲▲
 ▲   ▲
▲▲▲ ▲▲▲
`
	if printTriforce(3) != expected {
		t.Error("Bad result for 3")
	}
}

func TestTriForce5(t *testing.T) {
	expected := `
     ▲
    ▲▲▲
   ▲▲▲▲▲
  ▲     ▲
 ▲▲▲   ▲▲▲
▲▲▲▲▲ ▲▲▲▲▲
`
	if printTriforce(5) != expected {
		t.Error("Bad result for 5")
	}
}

func TestTriForce4(t *testing.T) {
	expected := `cannot use even number - only odd`
	if printTriforce(4) != expected {
		t.Error("Bad result for even number")
	}
}
