package main

import (
	"regexp"
	"testing"
)

func TestFB(t *testing.T) {
	result := "z"
	want := regexp.MustCompile(result)
	msg := Fizzbuzz()
	if !want.MatchString(msg) {
		t.Fatalf(`Fizzbuzz give us : %q - and we want %v`, msg, want)
	}
}
