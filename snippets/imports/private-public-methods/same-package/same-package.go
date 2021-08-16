package main

import "fmt"

func privateFuncSLSP() {
	fmt.Println("hello, private, same level, same package")
}

func PublicFuncSLSP() {
	fmt.Println("hello, public, same level, same package")
}
