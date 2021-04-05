package main

import (
	"regexp"
	"testing"
)

func TestFB2(t *testing.T) {
	want := "1 2"
	msg := Fizzbuzz(2)
	if want != msg {
		t.Fatalf(`Fizzbuzz give us : %q - and we want %v`, msg, want)
	}
}

func TestFB16(t *testing.T) {
	result := "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16"
	want := regexp.MustCompile(result)
	msg := Fizzbuzz(16)
	if !want.MatchString(msg) {
		t.Fatalf(`Fizzbuzz give us : %q - and we want %v`, msg, want)
	}
}

func BenchmarkFizz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fizzbuzz(16)
	}
}

// go test -bench=.
