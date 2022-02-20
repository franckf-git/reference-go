package main

import "fmt"
import "os"

func main() {
	fmt.Fprintln(os.Stdout, "standard message")
	fmt.Fprintln(os.Stderr, "err message")
}

// go run stdout-stderr.go >> stdout.log 2>> stderr.log
// go run stdout-stderr.go >> combined.log 2>&1
