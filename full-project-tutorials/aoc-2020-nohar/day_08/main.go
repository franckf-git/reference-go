package main

import (
	"fmt"
	"regexp"
	"strconv"

	"gitlab.com/neuware/aoc-2020/utils"
)

// OpType is the type of an opcode
type OpType uint8

const (
	OpNop OpType = iota // nop (= 0)
	OpJmp               // jmp (= 1)
	OpAcc               // acc (= 2)
)

// Opcode is a program instruction
type Opcode struct {
	Op  OpType
	Arg int
}

// Accepted string instructions
var opFmt = regexp.MustCompile(`^(jmp|acc|nop) ((\+|-)\d+)$`)

// load parses an input program into a sequence of opcodes
func load(input []string) []Opcode {
	program := make([]Opcode, len(input))
	for i, stmt := range input {
		groups := opFmt.FindStringSubmatch(stmt)
		if groups == nil {
			panic(fmt.Errorf("invalid statement at l.%d: %s", i, stmt))
		}

		var c Opcode
		c.Arg, _ = strconv.Atoi(groups[2])

		switch groups[1] {
		case "acc":
			c.Op = OpAcc
		case "jmp":
			c.Op = OpJmp
		case "nop":
			c.Op = OpNop
		}
		program[i] = c
	}
	return program
}

// eval runs the program, starting from given execution state ('pc' and 'acc').
// The footprint slice is used to detect loops by tracking instructions that
// have already been executed.
//
// Returns (acc, ok), where ok is true when the program terminates normally.
//
func eval(program []Opcode, footprint []bool, pc int, acc int) (int, bool) {
	for pc < len(program) && !footprint[pc] {
		footprint[pc] = true
		switch o := program[pc]; o.Op {
		case OpNop:
			pc++
		case OpAcc:
			acc += o.Arg
			pc++
		case OpJmp:
			pc += o.Arg
		}
	}
	return acc, pc == len(program)
}

// evalBacktrack runs the program, starting from given execution state ('pc'
// and 'acc'). When the a loop is detected, it backtracks to the last 'nop' or
// 'jmp' that was executed, flips it, and evaluates the rest of the resulting
// program.
//
// See 'eval' for arguments and return values.
//
func evalBacktrack(program []Opcode, footprint []bool, pc int, acc int) (int, bool) {
	for pc < len(program) && !footprint[pc] {
		footprint[pc] = true
		switch o := program[pc]; o.Op {
		case OpNop:
			if result, ok := evalBacktrack(program, footprint, pc+1, acc); ok {
				return result, ok
			}
			return eval(program, footprint, pc+o.Arg, acc)
		case OpJmp:
			if result, ok := evalBacktrack(program, footprint, pc+o.Arg, acc); ok {
				return result, ok
			}
			return eval(program, footprint, pc+1, acc)
		case OpAcc:
			acc += o.Arg
			pc++

		}
	}
	return acc, pc == len(program)
}

// run the program until it terminates or a loop is detected.
func run(program []Opcode) int {
	footprint := make([]bool, len(program))
	result, _ := eval(program, footprint, 0, 0)
	return result
}

// patch uses a recursive backtracking approach to run the program
// until it terminates. It is assumed that only one opcode is
// corrupted in the program.
func patch(program []Opcode) int {
	footprint := make([]bool, len(program))
	result, ok := evalBacktrack(program, footprint, 0, 0)
	if !ok {
		panic("couldn't find a successful patch")
	}
	return result
}

// bruteForce patches the program using a bruteforce approach
func bruteForce(program []Opcode) int {
	allFalse := make([]bool, len(program))
	footprint := make([]bool, len(program))
	for i, o := range program {
		switch o.Op {
		case OpAcc:
			continue
		case OpNop, OpJmp:
			program[i].Op ^= OpJmp
		}
		result, ok := eval(program, footprint, 0, 0)
		program[i].Op ^= OpJmp
		if ok {
			return result
		}
		copy(footprint, allFalse)
	}
	panic("couldn't find a successful patch")
}

func main() {
	program := load(utils.ReadLines())

	result := run(program)
	fmt.Println("Part 1:", result)

	result = patch(program)
	fmt.Println("Part 2:", result)

	result = bruteForce(program)
	fmt.Println("Brute force:", result)
}
