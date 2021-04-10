package main

import (
	"regexp"
	"testing"
)

func TestFizzBuzzGroup(t *testing.T) {
	t.Run("with 2", func(t *testing.T) {
		want := "1 2"
		got := Fizzbuzz(2)
		if want != got {
			t.Errorf("got - %v - want - %v -", got, want)
		}
	})
	t.Run("with 16", func(t *testing.T) {
		result := "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16"
		want := regexp.MustCompile(result)
		got := Fizzbuzz(16)
		if !want.MatchString(got) {
			t.Errorf("got - %v - want - %v -", got, want)
		}
	})
}
