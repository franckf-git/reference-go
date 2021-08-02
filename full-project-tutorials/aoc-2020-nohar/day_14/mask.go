package main

import (
	"math/bits"
)

type Mask struct {
	zeros  uint64
	ones   uint64
	floats uint64
}

func ParseMask(input string) Mask {
	if len(input) != 36 {
		panic("Invalid mask length")
	}

	m := Mask{}

	for i, c := range input {
		switch c {
		case '0':
			m.zeros |= 1 << (35 - i)
		case '1':
			m.ones |= 1 << (35 - i)
		case 'X':
			m.floats |= 1 << (35 - i)
		}
	}
	return m
}

func (m Mask) Apply(n uint64) uint64 {
	return n&^m.zeros | m.ones
}

func (m Mask) IterAddr(n uint64, callback func(uint64)) {
	base := n&m.zeros | m.ones
	var i uint64
	for {
		callback(base | i)
		i = (i | ^m.floats + 1) & m.floats
		if i == 0 {
			return
		}
	}
}

func (m Mask) Generate(n uint64) []uint64 {
	res := make([]uint64, 0, 1<<bits.OnesCount64(m.floats))
	m.IterAddr(n, func(g uint64) {
		res = append(res, g)
	})
	return res
}
