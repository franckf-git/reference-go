package main

import (
	"reflect"
	"testing"
)

func TestProcessNumEven(t *testing.T) {
	var want int = 4
	var msg int = ProcessNum(8)
	if want != msg {
		t.Fatalf(`test give us : %v - and we want %v`, msg, want)
	}
}

func TestProcessNumOdd(t *testing.T) {
	var want int = 10
	var msg int = ProcessNum(3)
	if want != msg {
		t.Fatalf(`test give us : %v - and we want %v`, msg, want)
	}
}

func TestCollatzArray(t *testing.T) {
	want := []int{8, 4, 2, 1}
	msg := CollatzProcess(8)
	if !reflect.DeepEqual(want, msg) {
		t.Fatalf(`test give us : %v - and we want %v`, msg, want)
	}
}
