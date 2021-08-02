package main

const DefaultThreadLength = 20

// A frame is a stackable state: it points to a subrule in an alternate
// sequence in a rule.
type Frame struct{ Rule, Seq, Sub int }

// A thread is an instantaneous state of the matcher
type Thread struct {
	stack []Frame
}

// NewThread creates a new thread, initialized to the first subrule of the
// first alternative of rule 0.
func NewThread() *Thread {
	return &Thread{make([]Frame, 1, DefaultThreadLength)}
}

// Clone creates a copy of the thread
func (t Thread) Clone() *Thread {
	clone := make([]Frame, len(t.stack), max(len(t.stack), DefaultThreadLength))
	copy(clone, t.stack)
	return &Thread{clone}
}

// Frame returns this thread's current frame
func (t *Thread) Frame() Frame {
	return t.stack[len(t.stack)-1]
}

// Set pushes a frame on top of this thread's stack
func (t *Thread) Set(f Frame) Frame {
	t.stack = append(t.stack, f)
	return f
}

// pop the top of the stack and return it.
// panic if the thread's stack is empty
func (t *Thread) pop() Frame {
	if t.Done() {
		panic("Thread stack empty")
	}
	last := len(t.stack) - 1
	f := t.stack[last]
	t.stack = t.stack[:last]
	return f
}

// Done returns True if the thread is finished
func (t Thread) Done() bool {
	return len(t.stack) == 0
}

// Match checks that given char matches the terminal symbol in the
// thread's current frame.
// If it matches, the thread is moved forward one step.
// Returns true if the character matches.
func (t *Thread) Match(char byte, grammar Grammar) bool {
	f := t.pop()
	r := grammar[f.Rule]
	if r.Terminal != char {
		return false
	}

	// Move thread to next step
	for !t.Done() {
		f = t.pop()
		r = grammar[f.Rule]
		f.Sub++
		if f.Sub < len(r.Sequences[f.Seq]) {
			// If the current rule needs to move forward,
			// update the current frame
			t.Set(f)
			return true
		}
	}
	// The thread has successfully terminated
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
