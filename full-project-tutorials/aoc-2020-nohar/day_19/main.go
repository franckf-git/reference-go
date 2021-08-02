package main

import (
	"fmt"

	"gitlab.com/neuware/aoc-2020/utils"
)

const MaxThreads = 300

func main() {
	grammar, input := parseInput(utils.ReadLines())
	fmt.Println("Part 1:", countMatches(input, grammar))
	grammar = patch(grammar)
	fmt.Println("Part 2:", countMatches(input, grammar))
}

// Patch grammar for part 2
func patch(grammar Grammar) Grammar {
	grammar[8] = Rule{Sequences: []Sequence{{42}, {42, 8}}}
	grammar[11] = Rule{Sequences: []Sequence{{42, 31}, {42, 11, 31}}}
	return grammar
}

// Count input strings that match the grammar
func countMatches(input []string, grammar []Rule) (count int) {
	for _, in := range input {
		if match(in, grammar) {
			count++
		}
	}
	return
}

// Match an input string against a grammar
func match(input string, grammar Grammar) bool {
	threads := make([]*Thread, 1, MaxThreads)
	threads[0] = NewThread()
	for _, c := range []byte(input) {
		// Create forks and prune terminated threads
		threads = expand(threads, grammar)
		// Move all matching frames forward, discard others
		threads = step(threads, grammar, c)
		if len(threads) == 0 {
			return false
		}
	}
	// Check if there is at least one terminated thread
	for _, t := range threads {
		if t.Done() {
			return true
		}
	}
	return false
}

// Expand the current threads until they all point to a terminal rule.
// If the threads are done, prune them away.
// If a rule with several alternatives is found during expansion, fork the
// corresponding threads.
func expand(threads []*Thread, grammar Grammar) []*Thread {
	expanded := threads[:0]
	for t := 0; t < len(threads); t++ {
		thread := threads[t]
		if thread.Done() {
			// If we reach here, then we have a terminated thread whereas
			// there are still chars to match in the input string,
			// discard this thread.
			continue
		}
		frame := thread.Frame()
		rule := grammar[frame.Rule]
		// Expand the thread until it points to a terminal rule
		for rule.Terminal == 0 {
			idx := rule.Sequences[frame.Seq][frame.Sub]
			rule = grammar[idx]
			for i := 1; i < len(rule.Sequences); i++ {
				// create a new thread with the alternative path
				// it'll be expanded in a future iteration
				fork := thread.Clone()
				fork.Set(Frame{Rule: idx, Seq: i})
				threads = append(threads, fork)
			}
			// keep expanding the current thread
			frame = thread.Set(Frame{Rule: idx})
		}
		expanded = append(expanded, thread)
	}
	return expanded
}

// Match all current threads against the current input char.
// Prune away threads that failed.
func step(threads []*Thread, grammar Grammar, c byte) []*Thread {
	next := threads[:0]
	for _, thread := range threads {
		if thread.Match(c, grammar) {
			next = append(next, thread)
		}
	}
	return next
}
