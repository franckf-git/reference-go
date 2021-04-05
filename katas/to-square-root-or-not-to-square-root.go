package kata
import "math"
func SquareOrSquareRoot(arr []int) []int{
  result := []int{}
  for _,value := range arr {
    square := math.Sqrt(float64(value))
    // check if float square is an int
    if square == float64(int64(square)) {
		result = append(result, int(square))
	  } else {
		result = append(result, value*value)
    }
  }
  return result
}

// top

package kata

import "math"

func SquareOrSquareRoot(arr []int) []int {
  results := make([]int, len(arr))

  for i, x := range arr {
    sqrt := math.Sqrt(float64(x))

    if sqrt == math.Floor(sqrt) {
      results[i] = int(sqrt)
    } else {
      results[i] = x * x
    }
  }

  return results
}
