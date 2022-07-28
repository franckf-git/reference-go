package main

import (
	"math"
	"testing"
)

func minMoves(disks int) (moves int) {
	// 2^n-1
	moves = int(math.Pow(2.0, float64(disks)))
	moves--
	return
}

func TestMinMoves(t *testing.T) {
	want := 127
	have := minMoves(7)
	if want != have {
		t.Errorf("want %v - have %v", want, have)
	}
}

func TestHanoi(t *testing.T) {
	towerA := []int{7, 6, 5, 4, 3, 2, 1}
	towerB := []int{}
	towerC := []int{}
	var moving int
	move(7, towerA, towerC, towerB, &moving)
	have := moving
	want := minMoves(7)
	if want != have {
		t.Errorf("want %v - have %v", want, have)
	}
}
