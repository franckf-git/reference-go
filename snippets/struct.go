package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {

	//var c Circle
	//c := new(Circle)
	//c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5}

	fmt.Println(circleArea(&c))
	fmt.Println(c.x, c.y, c.r)
}
