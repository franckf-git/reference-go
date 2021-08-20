package main

import "fmt"

type Creature struct {
	Species string
}

func (c *Creature) Reset() {
	c.Species = ""
}

func (c Creature) ResetNoPoit() {
	c.Species = ""
}

func main() {
	var creature Creature = Creature{Species: "shark"}
	fmt.Printf("1) %+v\n", creature) // shark
	creature.ResetNoPoit()
	fmt.Printf("2) %+v\n", creature) // shark
	creature.Reset()
	fmt.Printf("2) %+v\n", creature) // ""
}
