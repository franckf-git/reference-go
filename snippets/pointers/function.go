package main

import (
    "fmt"
)

type Human struct {
    name string
    age int
    place string
}

func f(h *Human) {
    fmt.Println("The user", (*h).name, "is", (*h).age, "years old and he is from", (*h).place)
}

func main() {
    john := Human{"John", 36, "Las Vegas"}
    f(&john) // The user John is 36 years old and he is from Las Vegas
}

// It is a very efficient way to pass large objects to function
// Be careful when dereferencing a struct. If you use it like *structname.field1 then it will throw an error. The correct way to do it is (*structname).field1.
// Using pointers inside functions makes the value mutable
