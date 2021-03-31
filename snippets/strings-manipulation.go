package main

import "fmt"

func main() {
	abc := "abcdefg"
	// on peut utiliser les positions mais cela transforme en byte uint8
	fmt.Printf("%T\n", abc[0])
	// premier caractere
	fmt.Print(string(abc[0]))
	fmt.Print("\n")
	// dernier caractere
	fmt.Print(string(abc[len(abc)-1]))
	fmt.Print("\n")
	// cinquieme caractere
	fmt.Print(string(abc[4]))
	fmt.Print("\n")
	// tout sauf le premier
	fmt.Print(string(abc[1:]))
	fmt.Print("\n")
	// les deux premiers
	fmt.Print(string(abc[:2]))
	fmt.Print("\n")
	// au centre cde
	fmt.Print(string(abc[2:5]))
	fmt.Print("\n")
}

