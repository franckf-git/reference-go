package kata

func Summation(n int) int {
	result := 0
	for i := 0; i <= n; i++ {
		result = result + i
	}
	return result
}

// top

package kata

func Summation(n int) int {
    return n * (n + 1) / 2
}