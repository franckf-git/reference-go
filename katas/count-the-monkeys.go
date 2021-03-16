package kata

func monkeyCount(n int) []int {
  var result = []int{}
  for i := 1; i <= n; i++ {
		result = append(result, i)
	}
  return result
}

// top

package kata

func monkeyCount(n int) []int {
  result := make([]int, n)
  for i := range result {
    result[i] = i + 1
  }
  return result
}
