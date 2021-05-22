package main

import (
    "fmt"
    "strconv"
    "log"
)

// type Stringer interface {
//    String() string
// }

// Declare a Book type which satisfies the fmt.Stringer interface.
type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Declare a Count type which satisfies the fmt.Stringer interface.
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}

// Declare a WriteLog() function which takes any object that satisfies
// the fmt.Stringer interface as a parameter.
func WriteLog(s fmt.Stringer) {
    log.Println(s.String())
}

func main() {
    // Initialize a Count object and pass it to WriteLog().
    book := Book{"Alice in Wonderland", "Lewis Carrol"}
    WriteLog(book)

    // Initialize a Count object and pass it to WriteLog().
    count := Count(3)
    WriteLog(count)
}

// We can use WriteLog on both types (Book and Count) because they both have a
// String method who satisfy the Stringer interface
// add security and allow multi types params

                                    // why ?
// - To help reduce duplication or boilerplate code. interfaces-boilerplate.go
// - To make it easier to use mocks instead of real objects in unit tests. interfaces-unit-mocking.go
// - As an architectural tool, to help enforce decoupling between parts of your codebase. interfaces-architectural.go

