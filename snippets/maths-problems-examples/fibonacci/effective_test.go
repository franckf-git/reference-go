package main

import "testing"

func TestGetValEff(t *testing.T) {
	for i := 0; i < 1000; i++ { // running it a 1000 times
		if Effective(30) != 832040 {
			t.Error("Incorrect!")
		}
	}
}
