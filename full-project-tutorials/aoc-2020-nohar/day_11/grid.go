package main

type Cell uint8

const (
	Floor Cell = iota
	Empty
	Occupied
)

type Grid struct {
	cells  []Cell
	width  int
	height int
}

func NewGrid(width, height int) *Grid {
	return &Grid{
		make([]Cell, width*height),
		width,
		height,
	}
}

func NewGridFromInput(input []string) *Grid {
	g := NewGrid(len(input[0]), len(input))
	for y, row := range input {
		for x, cell := range row {
			switch cell {
			case 'L':
				g.Set(x, y, Empty)
			case '#':
				g.Set(x, y, Occupied)
			}
		}
	}
	return g
}

func (g *Grid) At(x, y int) Cell {
	return g.cells[y*g.width+x]
}

func (g *Grid) Set(x, y int, c Cell) {
	g.cells[y*g.width+x] = c
}

func (g *Grid) Width() int {
	return g.width
}

func (g *Grid) Height() int {
	return g.height
}

func (g *Grid) Count(c Cell) int {
	var count int
	for _, cell := range g.cells {
		if cell == c {
			count++
		}
	}
	return count
}

func (g *Grid) ToSlice() []string {
	res := make([]string, g.height)
	for y := 0; y < g.height; y++ {
		row := make([]byte, g.width)
		for x := 0; x < g.width; x++ {
			switch g.At(x, y) {
			case Floor:
				row[x] = '.'
			case Empty:
				row[x] = 'L'
			case Occupied:
				row[x] = '#'
			}
		}
		res[y] = string(row)
	}
	return res
}
