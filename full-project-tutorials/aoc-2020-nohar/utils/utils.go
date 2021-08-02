package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mustOpen() *os.File {
	path := os.Getenv("INPUT_FILE")
	if path == "" {
		if len(os.Args) != 2 {
			Fatal(
				"Usage:", os.Args[0], "<input file>",
				"\nOr set INPUT_FILE env variable",
			)
		}
		path = os.Args[1]
	}
	f, err := os.Open(path)
	if err != nil {
		Fatal("Couldn't open file:", err)
	}
	return f
}

// ReadLines opens the input file and reads its lines.
func ReadLines() []string {
	f := mustOpen()
	defer f.Close()
	s := bufio.NewScanner(f)
	var input []string
	for s.Scan() {
		input = append(input, s.Text())
	}
	if len(input) == 0 {
		Fatal("Empty input")
	}
	return input
}

// ReadInts opens the input file and reads its lines as ints.
func ReadInts() []int {
	f := mustOpen()
	defer f.Close()
	s := bufio.NewScanner(f)
	var ints []int
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			Fatal(err)
		}
		ints = append(ints, i)
	}
	if len(ints) == 0 {
		Fatal("Empty input")
	}
	return ints
}

// Fatal exits the program after displaying the arguments on stderr
func Fatal(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

// Fatalf exits the program after displaying a formatted message on stderr
func Fatalf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
