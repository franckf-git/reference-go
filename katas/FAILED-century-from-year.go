// top

package kata

func century(year int) int {
	return (year + 99) / 100
}

// ---

package kata

func century(year int) int {
	// Finish this :)
	if year%100 != 0 {
		year += 100
	}
	return year / 100
}
