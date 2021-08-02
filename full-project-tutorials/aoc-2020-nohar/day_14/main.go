package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gitlab.com/neuware/aoc-2020/utils"
)

type ExecFunc func(map[uint64]uint64, Mask, uint64, uint64)

func run(program []string, exec ExecFunc) uint64 {
	var mask Mask
	mem := make(map[uint64]uint64)
	for _, line := range program {
		if strings.HasPrefix(line, "mask") {
			mask = ParseMask(line[len("mask = "):])
			continue
		}
		addr, arg := parseCmd(line)
		exec(mem, mask, addr, arg)
	}
	var sum uint64
	for _, v := range mem {
		sum += v
	}
	return sum
}

var cmdFormat = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

func parseCmd(line string) (addr, arg uint64) {
	groups := cmdFormat.FindStringSubmatch(line)
	if groups == nil {
		panic(fmt.Errorf("Malformed command %q", line))
	}
	addr, _ = strconv.ParseUint(groups[1], 10, 64)
	arg, _ = strconv.ParseUint(groups[2], 10, 64)
	return
}

func part1(mem map[uint64]uint64, mask Mask, addr, arg uint64) {
	mem[addr] = mask.Apply(arg)
}

func part2(mem map[uint64]uint64, mask Mask, addr, arg uint64) {
	mask.IterAddr(addr, func(x uint64) {
		mem[x] = arg
	})
}

func main() {
	program := utils.ReadLines()
	fmt.Println("Part 1:", run(program, part1))
	fmt.Println("Part 2:", run(program, part2))
}
