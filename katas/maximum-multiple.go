package kata
import "math"

func MaxMultiple(d, b int) int {
	var result int
	i := b
	for i > 0 {
		comp := float64(i) / float64(d)
		if comp == math.Floor(comp) {
			result = i
			break
		}
		i = i - 1
	}
	return result
}

// top

package kata

func MaxMultiple(d, b int) int {
    return b - b % d
}

// ---

package kata

func MaxMultiple(d, b int) int {
    return (b/d)*d
}
