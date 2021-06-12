package main

import (
    "fmt"
)

// declare interface
type Bird interface{
    fly()
}

type B struct{
    name string
}

// implement it
func (b B)fly() {
    fmt.Println("Flying...")
}

func main() {
    var a Bird = B{"Peacock"}
    b := &a
    fmt.Println(b)    // 0x40c138
    fmt.Println(*b)   // {Peacock}
}

// Here “a” is a struct of the type Bird which is then used for an interface type as you can see. This is polymorphism in action. Go allows polymorphism using interfaces. So, you can see pointers to a struct or an interface is an essential tool in Go.
