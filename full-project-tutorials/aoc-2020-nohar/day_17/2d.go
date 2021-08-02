package main

type Cell2D struct{ X, Y int }

func (c Cell2D) IterNeighborhood(yield func(Cell2D)) {
	for xx := c.X - 1; xx <= c.X+1; xx++ {
		for yy := c.Y - 1; yy <= c.Y+1; yy++ {
			yield(Cell2D{xx, yy})
		}
	}
}

type World2D map[Cell2D]bool

func World2DFromInput(input []string) World2D {
	w := make(World2D)
	for y, line := range input {
		for x, c := range line {
			if c == '#' {
				w.Add(Cell2D{x, y})
			}
		}
	}
	return w
}

func (w World2D) Add(c Cell2D) {
	w[c] = true
}

func (w World2D) Kill(c Cell2D) {
	delete(w, c)
}

func (w World2D) Step() int {
	counter := make(map[Cell2D]int)
	for c := range w {
		c.IterNeighborhood(func(n Cell2D) {
			counter[n]++
		})
	}
	for c, count := range counter {
		if count == 3 {
			w.Add(c)
		} else if count != 4 {
			w.Kill(c)
		}
	}
	return len(w)
}
