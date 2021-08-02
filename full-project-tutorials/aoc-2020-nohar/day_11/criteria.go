package main

func CanTakeSeatPart1(g *Grid, x, y int) bool {
	w, h := g.Width(), g.Height()
	for _, d := range directions {
		xx, yy := x+d.x, y+d.y
		if !inRange(xx, yy, w, h) {
			continue
		}
		if g.At(xx, yy) == Occupied {
			return false
		}
	}
	return true
}

func MustLeaveSeatPart1(g *Grid, x, y int) bool {
	w, h := g.Width(), g.Height()
	var count int
	for _, d := range directions {
		xx, yy := x+d.x, y+d.y
		if !inRange(xx, yy, w, h) {
			continue
		}
		if g.At(xx, yy) == Occupied {
			count++
		}
	}
	return count >= 4
}

func CanTakeSeatPart2(g *Grid, x, y int) bool {
	w, h := g.Width(), g.Height()
	for _, d := range directions {
		for xx, yy := x+d.x, y+d.y; inRange(xx, yy, w, h); xx, yy = xx+d.x, yy+d.y {
			c := g.At(xx, yy)
			if c == Occupied {
				return false
			}
			if c == Empty {
				break
			}
		}
	}
	return true
}

func MustLeaveSeatPart2(g *Grid, x, y int) bool {
	w, h := g.Width(), g.Height()
	var occupied int
	for _, d := range directions {
		for xx, yy := x+d.x, y+d.y; inRange(xx, yy, w, h); xx, yy = xx+d.x, yy+d.y {
			c := g.At(xx, yy)
			if c == Occupied {
				occupied++
			}
			if c != Floor {
				break
			}
		}
	}
	return occupied >= 5
}

var directions = []struct{ x, y int }{
	{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
}

func inRange(x, y, w, h int) bool {
	return 0 <= x && x < w && 0 <= y && y < h
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
