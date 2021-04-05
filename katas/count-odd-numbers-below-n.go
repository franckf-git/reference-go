package kata

func OddCount(n int) int{
var result []int
for i := 1; i < n; i++ {
      if i % 2 != 0{
result = append(result, i)
      }
	}
	return len(result)
}

// top

package kata

func OddCount(n int) int{
  return n/2
}

// ---

package kata

func OddCount(n int) int{
  return n >> 1;
}
