package main

import "fmt"

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})

	// Output
	// Ace of Hearts
}
