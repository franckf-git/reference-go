package main

type Cell4D struct{ X, Y, Z, W int }

func (c Cell4D) IterNeighborhood(yield func(Cell4D)) {
	for xx := c.X - 1; xx <= c.X+1; xx++ {
		for yy := c.Y - 1; yy <= c.Y+1; yy++ {
			for zz := c.Z - 1; zz <= c.Z+1; zz++ {
				for ww := c.W - 1; ww <= c.W+1; ww++ {
					yield(Cell4D{xx, yy, zz, ww})
				}
			}
		}
	}
}

type World4D map[Cell4D]bool

func World4DFromInput(input []string) World4D {
	w := make(World4D)
	for y, line := range input {
		for x, c := range line {
			if c == '#' {
				w.Add(Cell4D{x, y, 0, 0})
			}
		}
	}
	return w
}

func (w World4D) Add(c Cell4D) {
	w[c] = true
}

func (w World4D) Kill(c Cell4D) {
	delete(w, c)
}

func (w World4D) Step() int {
	counter := make(map[Cell4D]int)
	for c := range w {
		c.IterNeighborhood(func(n Cell4D) {
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
