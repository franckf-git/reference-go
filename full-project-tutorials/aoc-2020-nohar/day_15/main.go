package main

import "fmt"

const Stop = 2020

var Input = []int32{18, 11, 9, 0, 5, 1}

func countZeros(in []int32) {
	var count int
	for _, x := range in {
		if x == 0 {
			count++
		}
	}
	fmt.Printf(
		"Zeros: %d/%d (%.02f)\n",
		count, len(in), float32(count)/float32(len(in)),
	)
}

func MemoryGame(start []int32, end int32) int32 {
	var last, speak, pos int32 = 0, 0, 1
	spoken := make([]int32, end)
	for _, n := range start {
		last = n
		spoken[n] = pos
		pos++
	}
	for ; pos <= int32(end); pos++ {
		if spoken[last] != 0 {
			speak = pos - 1 - spoken[last]
		} else {
			speak = 0
		}
		spoken[last] = pos - 1
		last = speak
	}
	// countZeros(spoken)
	return last
}

func MemoryGameHashmap(start []int32, end int32) int32 {
	var last, speak, pos int32 = 0, 0, 1
	spoken := make(map[int32]int32)
	for _, n := range start {
		last = n
		spoken[n] = pos
		pos++
	}
	for ; pos <= int32(end); pos++ {
		if spoken[last] != 0 {
			speak = pos - 1 - spoken[last]
		} else {
			speak = 0
		}
		spoken[last] = pos - 1
		last = speak
	}
	return last
}

func main() {
	fmt.Println("Part 1:", MemoryGame(Input, 2020))
	fmt.Println("Part 2:", MemoryGame(Input, 30000000))
}
