package main

import (
	"fmt"
	"strconv"

	"gitlab.com/neuware/aoc-2020/utils"
)

type Direction struct{ X, Y int }

var (
	East  = Direction{1, 0}
	North = Direction{0, 1}
	West  = Direction{-1, 0}
	South = Direction{0, -1}

	Move = map[byte]Direction{
		'E': East,
		'N': North,
		'W': West,
		'S': South,
	}

	Cardinal = [4]Direction{East, North, West, South}
)

// Compute new rotation index
// * 'from' is a rotation index (0=East, 1=North, etc.)
// * 'angle' is given in counter-clockwise degrees (multiples of 90°)
//
// Returns the new rotation index
func rotate(from int, angle int) int {
	idx := (from + angle/90) % 4
	if idx < 0 {
		return idx + 4
	}
	return idx
}

// Rotate the waypoint around the ship.
// Returns the new position of the waypoint
func rotateWaypoint(x, y, angle int) (int, int) {
	// Get the cardinal direction from the rotation angle
	d := Cardinal[rotate(0, angle)]

	// Since we're working only with cardinal angles, we have this nifty
	// property: cosines and sines translate to vector coordinates.
	// e.g. cos(180) = "X of Cardinal[180/90]" = "X of West" = -1
	cos, sin := d.X, d.Y

	// Apply the rotation around the ship
	return x*cos - y*sin, x*sin + y*cos
}

func parseOrder(order string) (op byte, arg int) {
	var err error
	op = order[0]
	if arg, err = strconv.Atoi(order[1:]); err != nil {
		panic(fmt.Errorf("Bad order %q: %s", order, err))
	}
	return
}

func navigatePart1(input []string) (x, y int) {
	var orientation int // initial rotation index (0 = East)
	for _, instruction := range input {
		op, arg := parseOrder(instruction)
		switch op {
		case 'N', 'S', 'W', 'E':
			d := Move[op]
			x += arg * d.X
			y += arg * d.Y
		case 'F':
			d := Cardinal[orientation]
			x += arg * d.X
			y += arg * d.Y
		case 'L':
			orientation = rotate(orientation, arg)
		case 'R':
			orientation = rotate(orientation, -arg)
		}
	}
	return
}

func navigatePart2(input []string) (x, y int) {
	// Waypoint is initialized 10 East, 1 North
	wpX, wpY := 10, 1
	for _, instruction := range input {
		op, arg := parseOrder(instruction)
		switch op {
		case 'N', 'S', 'W', 'E':
			d := Move[op]
			wpX += arg * d.X
			wpY += arg * d.Y
		case 'F':
			x += arg * wpX
			y += arg * wpY
		case 'L':
			wpX, wpY = rotateWaypoint(wpX, wpY, arg)
		case 'R':
			wpX, wpY = rotateWaypoint(wpX, wpY, -arg)
		}
	}
	return
}

func manhattan(x, y int) int {
	return abs(x) + abs(y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	input := utils.ReadLines()

	x, y := navigatePart1(input)
	fmt.Println("Part 1:", manhattan(x, y))

	x, y = navigatePart2(input)
	fmt.Println("Part 2:", manhattan(x, y))
}
