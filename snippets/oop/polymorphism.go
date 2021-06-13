package main

import "fmt"

// declare interface
type Dog interface {
    Bark()
}

// declare struct
type Dalmatian struct {
    DogType string
}

// implement the interface
func (d Dalmatian) Bark() {
    fmt.Println("Dalmatian barking!!")
}

func MakeDogBark(d Dog) {
    d.Bark()
}

func main() {
    d := Dalmatian{"Jack"}
    MakeDogBark(d)                    // Dalmatian barking!!
}
