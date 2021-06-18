package main

import "testing"

func TestGetValEffMem(t *testing.T) {
	for i := 0; i < 5000000; i++ { // running it a 1000 times
		if EffectiveMem(30) != 832040 {
			t.Error("Incorrect!")
		}
	}
}
