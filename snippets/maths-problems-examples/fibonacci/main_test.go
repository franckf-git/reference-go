package main

import (
	"reflect"
	"testing"
)

func TestNormalArray(t *testing.T) {
	want := []int{1, 1}
	msg := Fibo(2)
	if !reflect.DeepEqual(want, msg) {
		t.Fatalf(`test give us : %v - and we want %v`, msg, want)
	}
}

func TestFullFibo(t *testing.T) {
	want := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}
	msg := Fibo(9)
	if !reflect.DeepEqual(want, msg) {
		t.Fatalf(`test give us : %v - and we want %v`, msg, want)
	}
}

func TestBinet(t *testing.T) {
	want := 12586269025
	msg := Binet(50)
	if want != msg {
		t.Fatalf(`test give us : %v - and we want %v`, msg, want)
	}
}
