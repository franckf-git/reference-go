package main

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/neuware/aoc-2020/utils"
)

func parseInput() (int, []int) {
	input := utils.ReadLines()
	now, err := strconv.Atoi(input[0])
	if err != nil {
		panic("invalid input")
	}
	buses := make([]int, len(input[1]))
	for i, bus := range strings.Split(input[1], ",") {
		buses[i], _ = strconv.Atoi(bus)
	}
	return now, buses
}

func part1(now int, buses []int) (int, int) {
	best := 0
	minDelay := now
	for _, bus := range buses {
		if bus == 0 {
			continue
		}
		delay := bus - (now % bus)
		if delay < minDelay {
			best, minDelay = bus, delay
		}
	}
	return best, minDelay
}

func part2(buses []int) int {
	timestamp, step := 0, 1
	for i, bus := range buses {
		if bus == 0 {
			// No constraint, move on to the next bus
			continue
		}

		// Look for the first timestamp that satisfies this bus' constraint
		// as well as all previous ones.
		for (timestamp+i)%bus != 0 {
			timestamp += step
		}

		// All constraints were satisfied up until now.
		// We know this won't happen before another (step * bus) cycles
		// because the buses are prime numbers, so we adjust the step
		// accordingly.
		step *= bus
	}
	return timestamp
}

func main() {
	now, buses := parseInput()
	best, delay := part1(now, buses)
	fmt.Println("Part 1:", best*delay)
	fmt.Println("Part 2:", part2(buses))
}
