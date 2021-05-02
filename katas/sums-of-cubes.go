package kata

func SumCubes(n int) int {
	result := 0
	for i := 0; i <= n; i++ {
		result = i*i*i + result
	}
	return result
}

// top (no loop O(n) ?)

package kata

func SumCubes(n int) int {
  n = n * (n + 1)
  return n * n / 4
}
