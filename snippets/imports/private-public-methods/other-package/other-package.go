package samelevel

import "fmt"

func privateFuncSLOP() {
	fmt.Println("hello, private, same level, other package")
}

func PublicFuncSLOP() {
	fmt.Println("hello, public, same level, other package")
}
