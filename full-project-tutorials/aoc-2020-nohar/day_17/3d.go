package main

type Cell3D struct{ X, Y, Z int }

func (c Cell3D) IterNeighborhood(yield func(Cell3D)) {
	for xx := c.X - 1; xx <= c.X+1; xx++ {
		for yy := c.Y - 1; yy <= c.Y+1; yy++ {
			for zz := c.Z - 1; zz <= c.Z+1; zz++ {
				yield(Cell3D{xx, yy, zz})
			}
		}
	}
}

type World3D map[Cell3D]bool

func World3DFromInput(input []string) World3D {
	w := make(World3D)
	for y, line := range input {
		for x, c := range line {
			if c == '#' {
				w.Add(Cell3D{x, y, 0})
			}
		}
	}
	return w
}

func (w World3D) Add(c Cell3D) {
	w[c] = true
}

func (w World3D) Kill(c Cell3D) {
	delete(w, c)
}

func (w World3D) Step() int {
	counter := make(map[Cell3D]int)
	for c := range w {
		c.IterNeighborhood(func(n Cell3D) {
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
