package main

import (
	"reflect"
	"testing"
)

func TestFlipH(t *testing.T) {
	base := Image{
		[]byte{0x0, 0x1, 0x2, 0x3},
		[]byte{0x4, 0x5, 0x6, 0x7},
		[]byte{0x8, 0x9, 0xa, 0xb},
		[]byte{0xc, 0xd, 0xe, 0xf},
	}
	expected := Image{
		[]byte{0x3, 0x2, 0x1, 0x0},
		[]byte{0x7, 0x6, 0x5, 0x4},
		[]byte{0xb, 0xa, 0x9, 0x8},
		[]byte{0xf, 0xe, 0xd, 0xc},
	}

	res := base.FlipH()
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Exp. %v\nGot  %v\n", expected, res)
	}
}

func TestFlipV(t *testing.T) {
	base := Image{
		[]byte{0x0, 0x1, 0x2, 0x3},
		[]byte{0x4, 0x5, 0x6, 0x7},
		[]byte{0x8, 0x9, 0xa, 0xb},
		[]byte{0xc, 0xd, 0xe, 0xf},
	}
	expected := Image{
		[]byte{0xc, 0xd, 0xe, 0xf},
		[]byte{0x8, 0x9, 0xa, 0xb},
		[]byte{0x4, 0x5, 0x6, 0x7},
		[]byte{0x0, 0x1, 0x2, 0x3},
	}

	res := base.FlipV()
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Exp. %v\nGot  %v\n", expected, res)
	}
}

func TestFlipD(t *testing.T) {
	base := Image{
		[]byte{0x0, 0x1, 0x2, 0x3},
		[]byte{0x4, 0x5, 0x6, 0x7},
		[]byte{0x8, 0x9, 0xa, 0xb},
		[]byte{0xc, 0xd, 0xe, 0xf},
	}
	expected := Image{
		[]byte{0x0, 0x4, 0x8, 0xc},
		[]byte{0x1, 0x5, 0x9, 0xd},
		[]byte{0x2, 0x6, 0xa, 0xe},
		[]byte{0x3, 0x7, 0xb, 0xf},
	}
	res := base.FlipD()
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Exp. %v\nGot  %v\n", expected, res)
	}
}

func TestRotateRight(t *testing.T) {
	base := Image{
		[]byte{0x0, 0x1, 0x2, 0x3},
		[]byte{0x4, 0x5, 0x6, 0x7},
		[]byte{0x8, 0x9, 0xa, 0xb},
		[]byte{0xc, 0xd, 0xe, 0xf},
	}
	expected := Image{
		[]byte{0xc, 0x8, 0x4, 0x0},
		[]byte{0xd, 0x9, 0x5, 0x1},
		[]byte{0xe, 0xa, 0x6, 0x2},
		[]byte{0xf, 0xb, 0x7, 0x3},
	}
	res := base.RotateRight()
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Exp. %v\nGot  %v\n", expected, res)
	}
}

func TestRotateLeft(t *testing.T) {
	base := Image{
		[]byte{0x0, 0x1, 0x2, 0x3},
		[]byte{0x4, 0x5, 0x6, 0x7},
		[]byte{0x8, 0x9, 0xa, 0xb},
		[]byte{0xc, 0xd, 0xe, 0xf},
	}
	expected := Image{
		[]byte{0x3, 0x7, 0xb, 0xf},
		[]byte{0x2, 0x6, 0xa, 0xe},
		[]byte{0x1, 0x5, 0x9, 0xd},
		[]byte{0x0, 0x4, 0x8, 0xc},
	}
	res := base.RotateLeft()
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Exp. %v\nGot  %v\n", expected, res)
	}
}
